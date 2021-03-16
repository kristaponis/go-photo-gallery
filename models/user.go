package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //gorm postgres driver
	"github.com/kristaponis/go-photo-gallery/helpers"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is used when no record in database is found.
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is used when passed ID is invalid.
	ErrInvalidID = errors.New("models: ID is invalid, must be > 0")

	// ErrInvalidPassword is used when passed email address is invalid.
	ErrInvalidPassword = errors.New("models: email address is invalid")

	// passPepper ads additional fixed string to user password (pepper).
	passwordPepper = "chili-pepper"
)

// UserDB interface is used to interact with the database.
// For single user queries it will look for a user with the provided ID.
// If the user is found, then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info. Any error but ErrNotFound should be 500 error.
type UserDB interface {
	// Methods for querying for single users
	ByID(id uint) (*User, error)
	ByEmail(e string) (*User, error)
	ByRemember(t string) (*User, error)

	// Methods for altering users
	Create(u *User) error
	Update(u *User) error
	Delete(id uint) error

	// Method to close connection
	Close() error
}

// User holds template for user info to be inserted into database.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

// UserService hods gorm.DB object and performs database operations
// with its methods like ByID, Create and others.
type UserService struct {
	UserDB
}

type userValidator struct {
	UserDB
}

// This tests userGorm to match UserDB
var _ UserDB = &userGorm{}

// userGorm type
type userGorm struct {
	db   *gorm.DB
	hmac helpers.HMAC
}

// ByID will look for a user with the provided ID. If the user
// is found, then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info.
// Any error but ErrNotFound should be 500 error.
func (ug *userGorm) ByID(id uint) (*User, error) {
	var u *User
	err := ug.db.First(u).Where("id = ?", id).Error
	switch err {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// ByEmail will look for a user with the provided email. If the user
// is found, then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info.
// Any error but ErrNotFound should be 500 error.
func (ug *userGorm) ByEmail(e string) (*User, error) {
	var u User
	err := ug.db.Where("email = ?", e).First(&u).Error
	switch err {
	case nil:
		return &u, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// ByRemember will look for a user with the provided token and
// handle the hashing of the token. If the user is found,
// then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info.
// Any error but ErrNotFound should be 500 error.
func (ug *userGorm) ByRemember(t string) (*User, error) {
	var u User
	rememberHash := ug.hmac.HashString(t)
	err := ug.db.Where("remember_hash = ?", rememberHash).First(&u).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	case nil:
		return &u, nil
	default:
		return nil, err
	}
}

// Create will create the provided user, auto fill data and
// insert this info into database.
func (ug *userGorm) Create(u *User) error {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(u.Password+passwordPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashpass)
	u.Password = ""
	if u.Remember == "" {
		token, err := helpers.RememberToken()
		if err != nil {
			return err
		}
		u.Remember = token
	}
	u.RememberHash = ug.hmac.HashString(u.Remember)
	// us.db.AutoMigrate(&u)
	return ug.db.Create(&u).Error
}

// Update will update yhe provided user. It will rewrite the user
// with the new data.
func (ug *userGorm) Update(u *User) error {
	if u.Remember != "" {
		u.RememberHash = ug.hmac.HashString(u.Remember)
	}
	return ug.db.Save(u).Error
}

// Delete will delete the user with the provided id.
func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	uID := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&uID).Error
}

// Close closes UserService database connection.
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

// Authenticate is used to check the provided email e and password p.
// If they are correct, return the user, otherwise return error.
func (us *UserService) Authenticate(e string, p string) (*User, error) {
	u, err := us.ByEmail(e)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(p+passwordPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}
	return u, nil
}

// UserServiceConn opens connection with database.
func UserServiceConn(conn string) (*UserService, error) {
	ug, err := newUserGorm(conn)
	if err != nil {
		return nil, err
	}
	return &UserService{
		UserDB: userValidator{
			UserDB: ug,
		},
	}, nil
}

// NewUserGorm func
func newUserGorm(conn string) (*userGorm, error) {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := helpers.NewHMAC("temp-secret-key")
	return &userGorm{
		db:   db,
		hmac: hmac,
	}, nil
}

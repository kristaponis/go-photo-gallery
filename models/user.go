package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is used when no record in database is found.
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is used when passed ID is invalid.
	ErrInvalidID = errors.New("models: ID is invalid, must be > 0")

	// ErrInvalidPassword is used when passed email address is invalid.
	ErrInvalidPassword = errors.New("models: email adress is invalid")

	// passPepper ads additional fixed string to user password (pepper).
	passwordPepper = "chili-pepper"
)

// User holds template for user info to be inserted into database.
type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

// UserService hods gorm.DB object and performs database operations
// with its methods like ByID, Create and others.
type UserService struct {
	db *gorm.DB
}

// ByID will look for a user with the provided ID. If the user
// is found, then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info.
// Any error but ErrNotFound should be 500 error.
func (us *UserService) ByID(id uint) (*User, error) {
	var u *User
	err := us.db.First(u).Where("id = ?", id).Error
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
func (us *UserService) ByEmail(e string) (*User, error) {
	var u User
	err := us.db.First(&u).Where("email = ?", e).Error
	switch err {
	case nil:
		return &u, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create will create the provided user, auto fill data and
// insert this info into database.
func (us *UserService) Create(u *User) error {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(u.Password+passwordPepper), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashpass)
	u.Password = ""
	us.db.AutoMigrate(&u)
	return us.db.Create(&u).Error
}

// Update will update yhe provided user. It will rewrite the user
// with the new data.
func (us *UserService) Update(u *User) error {
	return us.db.Save(u).Error
}

// Delete will delete the user with the provided id.
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	userID := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&userID).Error
}

// Close closes UserService database connection.
func (us *UserService) Close() error {
	return us.db.Close()
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
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

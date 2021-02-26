package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound  = errors.New("models: resource not found")
	ErrInvalidID = errors.New("models: ID is invalid, must be > 0")
)

// User holds template for user info to be inserted into database.
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Pass     string `gorm:"-"`
	PassHash string `gorm:"not null"`
}

// UserService hods gorm.DB object and performs database operations
// with its methods like ByID, Create and others.
type UserService struct {
	db *gorm.DB
}

// ById will look for a user with the provided ID. If the user
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
	var u *User
	err := us.db.First(u).Where("email = ?", e).Error
	switch err {
	case nil:
		return u, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create will create the provided user, auto fill data and
// insert this info into database.
func (us *UserService) Create(u *User) error {
	return us.db.Create(u).Error
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

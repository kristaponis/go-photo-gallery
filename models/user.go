package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound = errors.New("models: resource not found")
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

type UserService struct {
	db *gorm.DB
}

// ById method will look for a user with the ID. If the user
// is found, then it will return the user and nil for the error.
// If the user is not found, it will return ErrNotFound error
// and nil for the user. If there is another error, it will return
// error with more info.
// Any error but ErrNotFound should be 500 error.
func (us *UserService) ByID(id uint) (*User, error) {
	var user *User
	err := us.db.First(user).Where("id = ?", id).Error
	switch err {
	case nil:
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create will create the provided user and auto fill data.
func (us *UserService) Create(u *User) error {
	return us.db.Create(u).Error
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

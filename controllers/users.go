package controllers

import (
	"net/http"

	"github.com/kristaponis/go-photo-gallery/views"
)

// Users type contains info about user
type Users struct {
	NewView *views.View
}

// New method renders the form used to create new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create method processes the sign up form when a new user tries
// to create a new account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All good!"))
}

// NewUsers generates new page from template with the form for signing up.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

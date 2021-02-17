package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/kristaponis/go-photo-gallery/views"
)

// SignupForm type represents new user signup form at /signup.
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Users type contains info about user.
type Users struct {
	NewView *views.View
}

// New renders the form used to create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create processes the sign up form when a new user tries
// to create a new account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var sf SignupForm
	if err := parseDecodeForm(r, &sf); err != nil {
		panic(err)
	}
}

// NewUsers generates new page from template with the form for signing up.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// parseDecodeForm parses passed form and decodes it with
// gorilla/schema package.
func parseDecodeForm(r *http.Request, dest interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	d := schema.NewDecoder()
	if err := d.Decode(dest, r.PostForm); err != nil {
		return err
	}
	return nil
}

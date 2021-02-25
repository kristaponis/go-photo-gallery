package controllers

import (
	"github.com/kristaponis/go-photo-gallery/models"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/kristaponis/go-photo-gallery/views"
)

// SignupForm type represents new user signup form at /signup.
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Users type contains info about user.
type Users struct {
	NewView     *views.View
	userService *models.UserService
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
	user := models.User{
		Name:  sf.Name,
		Email: sf.Email,
	}
	if err := u.userService.Create(&user); err != nil {
		http.Error(w, "Something wrong", http.StatusInternalServerError)
		panic(err)
	}
}

// NewUsers generates new page from template with the form for signing up.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "views/users/new.gohtml"),
		userService: us,
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

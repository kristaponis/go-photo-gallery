package controllers

import (
	"fmt"
	"net/http"

	"github.com/kristaponis/go-photo-gallery/models"

	"github.com/gorilla/schema"
	"github.com/kristaponis/go-photo-gallery/views"
)

// SignupForm represents new user signup form at /signup.
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// LoginForm represents existing user login form at /login.
type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Users type contains info about user.
type Users struct {
	NewView     *views.View
	LoginView   *views.View
	userService *models.UserService
}

// NewUser renders the form used to create a new user account.
//
// GET /signup
func (u *Users) NewUser(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// NewLogin renders the form used to log in for a existing user.
//
// GET /signup
func (u *Users) NewLogin(w http.ResponseWriter, r *http.Request) {
	if err := u.LoginView.Render(w, nil); err != nil {
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
		Name:     sf.Name,
		Email:    sf.Email,
		Password: sf.Password,
	}
	if err := u.userService.Create(&user); err != nil {
		http.Error(w, "Something wrong", http.StatusInternalServerError)
		panic(err)
	}
}

// Login is used to verify email address and password of the user,
// and log the user in if they are correct.
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var lf LoginForm
	if err := parseDecodeForm(r, &lf); err != nil {
		panic(err)
	}
	user, err := u.userService.Authenticate(lf.Email, lf.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email address")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password")
	case nil:
		fmt.Fprintln(w, user)
	default:
		http.Error(w, "Intermal server error - 500", http.StatusInternalServerError)
	}
}

// NewUsers generates new page from template with the form for signing up.
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "views/users/new.gohtml"),
		LoginView:   views.NewView("bootstrap", "views/users/login.gohtml"),
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

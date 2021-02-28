package controllers

import (
	"net/http"

	"github.com/kristaponis/go-photo-gallery/views"
)

// StaticPage type contains static pages.
type StaticPage struct {
	HomeView    *views.View
	ContactView *views.View
	UserView    *views.View
}

// NewStaticPage creates static pages views.
func NewStaticPage() *StaticPage {
	return &StaticPage{
		HomeView:    views.NewView("bootstrap", "views/static/home.gohtml"),
		ContactView: views.NewView("bootstrap", "views/static/contact.gohtml"),
		UserView:    views.NewView("bootstrap", "views/static/user.gohtml"),
	}
}

// Home renders main Home page.
func (sp *StaticPage) Home(w http.ResponseWriter, r *http.Request) {
	if err := sp.HomeView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Contact renders Contact page.
func (sp *StaticPage) Contact(w http.ResponseWriter, r *http.Request) {
	if err := sp.ContactView.Render(w, nil); err != nil {
		panic(err)
	}
}

// User renders Contact page after signup or login.
func (sp *StaticPage) User(w http.ResponseWriter, r *http.Request) {
	if err := sp.UserView.Render(w, nil); err != nil {
		panic(err)
	}
}

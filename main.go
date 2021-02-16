package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kristaponis/go-photo-gallery/controllers"
	"github.com/kristaponis/go-photo-gallery/views"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView := views.NewView("bootstrap", "views/home.gohtml")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactView := views.NewView("bootstrap", "views/contact.gohtml")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("<h3>404 Error - Not Found</h3>"))
	w.Write([]byte("<p><a href=\"/\">Go to Home Page</a></p>"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	r.HandleFunc("/contact", contactHandler).Methods(http.MethodGet)
	r.HandleFunc("/signup", controllers.NewUsers().New).Methods(http.MethodGet)
	r.HandleFunc("/signup", controllers.NewUsers().Create).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fmt.Println("Running server on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

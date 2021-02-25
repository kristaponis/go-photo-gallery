package main

import (
	"fmt"
	"github.com/kristaponis/go-photo-gallery/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kristaponis/go-photo-gallery/config"
	"github.com/kristaponis/go-photo-gallery/controllers"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("<h3>404 Error - Not Found</h3>"))
	w.Write([]byte("<p><a href=\"/\">Go to Home Page</a></p>"))
}

func main() {
	dns := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
		config.DBHOST, config.DBPORT, config.DBNAME)
	us, err := models.UserServiceConn(dns)
	if err != nil {
		panic(err)
	}
	ctrls := controllers.NewStaticPage()

	r := mux.NewRouter()

	r.HandleFunc("/", ctrls.Home).Methods(http.MethodGet)
	r.HandleFunc("/contact", ctrls.Contact).Methods(http.MethodGet)
	r.HandleFunc("/signup", controllers.NewUsers(us).New).Methods(http.MethodGet)
	r.HandleFunc("/signup", controllers.NewUsers(us).Create).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	fmt.Println("Running server on http://127.0.0.1:8080")
	http.ListenAndServe(config.PORT, r)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Hello</h1>"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	fmt.Println("Running server on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

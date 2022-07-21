package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to links, a bookmarking web app</h1>")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	fmt.Fprintf(w, "Hello %s", username)
}

func main() {
	log.Print("Starting links webserver")

	r := chi.NewRouter()
	r.Get("/", home)
	r.Get("/{username}", userHandler)

	http.ListenAndServe(":8080", r)

}

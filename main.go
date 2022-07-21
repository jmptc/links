package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "<h1>Welcome to links, a bookmarking web app</h1>")
    executeTemplate(w, filepath.Join("templates", "home.html"), nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	fmt.Fprintf(w, "Hello %s", username)
}

// executeTemplate parses and executes template 
func executeTemplate(w http.ResponseWriter, filepath string, data any) {
    t, err := template.ParseFiles(filepath)
    if err != nil {
        log.Print("error parsing template", err)
        http.Error(w, "error parsing template", http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, data)
    if err != nil {
        log.Print("error executing template", err)
        http.Error(w, "error executing template", http.StatusInternalServerError)
        return
    }
}

func main() {
	log.Print("Starting links webserver")

	r := chi.NewRouter()
    r.Use(middleware.Logger)

	r.Get("/", home)
	r.Get("/{username}", userHandler)

	http.ListenAndServe(":8080", r)

}

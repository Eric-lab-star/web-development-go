package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	r.Get("/contact", contact)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		notFound(w)
	})

	http.ListenAndServe("localhost:8080", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join("template", "home.gohtml")
	render(w, file)
}

func contact(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join("template", "contact.gohtml")
	render(w, file)
}

func faq(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join("template", "faq.gohtml")
	render(w, file)
}

func render(w http.ResponseWriter, file string) {

	w.Header().Set("Content-Type", "text/html; charset=utf8")
	tpl, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("error: %s", file)
		http.Error(w, "Parsing template error", http.StatusInternalServerError)
	}
	tpl.Execute(w, nil)
}
func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<h1>Page Not Found</h1>"))

}

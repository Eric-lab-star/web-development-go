package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Eric-lab-star/web-development/server/views"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func render(w http.ResponseWriter, filePath string) {
	tpl, err := views.Parse(filePath)
	if err != nil {
		log.Printf("error parsing template from file path: %s", filePath)
		http.Error(w, "Parsing template error", http.StatusInternalServerError)
	}
	tpl.Execute(w, nil)
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

func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<h1>Page Not Found</h1>"))
}
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	r.Get("/contact", contact)
	r.Get("/faq", faq)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		notFound(w)
	})
	log.Println("Server is listening to https://localhost:8080.")
	http.ListenAndServe("localhost:8080", r)
}

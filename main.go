package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	r.Get("/user/{userId}", profile)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		notFound(w)
	})
	apiRouter := chi.NewRouter()
	apiRouter.Get("/books", getBooks)
	r.Mount("/api", apiRouter)

	http.ListenAndServe("localhost:8080", r)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	render(w, "books")
}

func profile(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	render(w, "userId: "+userId)
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "home")
}

func render(w http.ResponseWriter, v string) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte(v))
}
func notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<h1>Page Not Found</h1>"))

}

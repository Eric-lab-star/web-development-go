package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/Eric-lab-star/web-development/server/controllers"
	"github.com/Eric-lab-star/web-development/server/views"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	tpl, err := views.Parse(filepath.Join("template", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("template", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))
	tpl, err = views.Parse(filepath.Join("template", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("template", "404.gohtml"))
	if err != nil {
		panic(err)
	}
	r.NotFound(controllers.StaticHandler(tpl))

	log.Println("Server is listening to https://localhost:8080.")
	http.ListenAndServe("localhost:8080", r)
}

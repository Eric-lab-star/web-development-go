package main

import (
	"fmt"
	"net/http"
)

func main() {

	Home := page{
		header: "<h1>Home page</h1>",
	}
	About := page{
		header: "<h1>About</h1>",
	}
	http.Handle("/", Home)
	http.Handle("/about", About)
	http.ListenAndServe("localhost:8080", nil)

}

type page struct {
	header string
}

func (page page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	print(page, w)
}

func print(page page, w http.ResponseWriter) {
	fmt.Fprint(w, page.header)
}

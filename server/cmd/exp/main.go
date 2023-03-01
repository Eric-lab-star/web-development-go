package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
		Bio  string
	}{
		Name: "Kim Kyungsub",
		Bio:  `<script>alert("hello world");</script>`,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err = t.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})
	fmt.Println("listening to server")
	http.ListenAndServe("localhost:8080", nil)

}

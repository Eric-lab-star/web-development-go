package main

import (
	"fmt"
	"net/http"
)

var PORT = "3000"

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Listening to server on http://localhost:3000")
	http.ListenAndServe("localhost:"+PORT, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to go server")
}

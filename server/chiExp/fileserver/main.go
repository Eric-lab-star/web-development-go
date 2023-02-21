// FileServer
// ===========
// This example demonstrates how to serve static files from your filesystem.
//
// Boot the server:
// ----------------
// $ go run main.go
//
// Client requests:
// ----------------
// $ curl http://localhost:3333/files/
// <pre>
// <a href="notes.txt">notes.txt</a>
// </pre>
//
// $ curl http://localhost:3333/files/notes.txt
// Notessszzz
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index handler
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hi"))
	// })

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "data"))
	staticDir := http.Dir(filepath.Join(workDir, "static"))
	client := http.Dir("/Users/kimkyungsub/go/src/github.com/Eric-lab-star/web-development/client/build")
	FileServer(r, "/files", filesDir)
	FileServer(r, "/exam", staticDir)
	FileServer(r, "/client", client)

	fmt.Println("sever running..")
	http.ListenAndServe("localhost:3333", r)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}
	//path는 '/'로 끝나도록 설정해준다.
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	//path에 '*'를 붙인다.
	path += "*"

	//변경된 path로 get요청
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {

		rctx := chi.RouteContext(r.Context())
		fmt.Println(rctx.RoutePattern())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fmt.Println(http.FileServer(root))
		//stripPrefix 하는 이유는 경로"/files"로 get요청이 왔을 때, /files를 url으로 부터 제거한다.
		//http.FileServer는 fileHandler 스트럭트 포인터를 반환하고 root의 파일에 접근할 수 있게한다.
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)

	})
}

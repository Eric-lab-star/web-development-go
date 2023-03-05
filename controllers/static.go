package controllers

import (
	"net/http"

	"github.com/Eric-lab-star/web-development/server/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

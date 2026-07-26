package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		files := []string{
			"./ui/html/404.go.tpl",
			"./ui/html/base.go.tpl",
		}

		w.WriteHeader(http.StatusNotFound)

		tmpl := template.Must(template.ParseFiles(files...))
		tmpl.Execute(w, nil)

		return
	}

	files := []string{
		"./ui/html/index.go.tpl",
		"./ui/html/base.go.tpl",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

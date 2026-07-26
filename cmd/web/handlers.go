package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/index.go.tpl",
		"./ui/html/base.go.tpl",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

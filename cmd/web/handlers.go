package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		files := []string{
			"./ui/html/404.html",
			"./ui/html/base.html",
		}

		w.WriteHeader(http.StatusNotFound)

		tmpl := template.Must(template.ParseFiles(files...))
		tmpl.Execute(w, nil)

		return
	}

	files := []string{
		"./ui/html/index.html",
		"./ui/html/base.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/contact.html",
		"./ui/html/base.html",
	}

	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, nil)
}

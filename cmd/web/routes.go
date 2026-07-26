package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)

	mux.HandleFunc("/contact", ContactHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

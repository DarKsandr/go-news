package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	//Application routes
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/contact", ContactHandler)
	mux.HandleFunc("/news", NewsHandler)
	mux.HandleFunc("/news/{id}", NewsDetailHandler)

	//Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//API
	mux.HandleFunc("GET /api/news", NewsGetApiHandler)

	return mux
}

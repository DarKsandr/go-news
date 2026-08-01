package main

import (
	_ "main/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

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
	mux.HandleFunc("/api/news", NewsGetApiHandler)

	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.Handler())

	return mux
}

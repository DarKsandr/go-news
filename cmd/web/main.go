package main

//go:generate swag init -g cmd/web/main.go -o cmd/web/docs

import (
	"log"
	"main/pkg"
	"net/http"
	"os"
)

// @title           Blueprint API
// @version         1.0
// @description     This is a sample RESTful API server.
// @host            localhost:8080
// @BasePath        /api
func main() {
	pkg.Init()

	port := os.Getenv("PORT")

	log.Println("Server is running on port:", port)

	mux := routes()

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}

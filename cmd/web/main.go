package main

import (
	"log"
	"main/pkg"
	"net/http"
	"os"
)

// @title           API
// @version         1.0
// @description     API server
// @BasePath        /api
func main() {
	pkg.Init()

	port := os.Getenv("PORT")

	log.Println("Server is running on port:", port)

	mux := routes()

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}

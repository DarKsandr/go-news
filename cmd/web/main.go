package main

import (
	"log"
	"main/pkg"
	"net/http"
	"os"
)

func main() {
	pkg.Init()

	port := os.Getenv("PORT")

	log.Println("Server is running on port:", port)

	mux := routes()

	log.Fatalln(http.ListenAndServe(":"+port, mux))
}

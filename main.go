package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aman123at/go-postgres/routes"
)

func main() {
	fmt.Println("Welcome to go with postgresql")

	router := routes.Router()

	fmt.Println("Starting server at 4000...")

	log.Fatal(http.ListenAndServe(":4000", router))
}

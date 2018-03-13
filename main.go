package main

import (
	"log"
	"net/http"
)

func main() {
	router := GetRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}

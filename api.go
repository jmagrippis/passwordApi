package main

import (
	"github.com/julienschmidt/httprouter"

	"log"
	"net/http"
)

// Defines the endpoints of the API
func main() {
	router := httprouter.New()
	router.GET("/", Welcome)
	router.GET("/generate/:amount", Generate)
	router.GET("/generate/:amount/safe", GenerateSafe)

	log.Fatal(http.ListenAndServe(":8090", router))
}

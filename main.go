package main

import (
	"./web"
	"log"
	"net/http"
)

/**
 *
 */
func main() {
	http.Handle("/", web.Routes())
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

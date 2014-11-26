package main

import (
	"./web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/{collection}/", web.Collection).Methods("GET")
	rtr.HandleFunc("/{collection}/{future}/", web.Future).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

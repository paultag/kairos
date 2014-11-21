package main

import (
	"./web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/foo/", web.Foo).Methods("GET")
	rtr.HandleFunc("/{path:.+}", web.Root).Methods("GET")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

package web

import (
	"../future"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Collection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["collection"]
	fmt.Fprintf(w, collection)
}

func Future(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	futureId := vars["future"]
	f := future.Get(futureId)
	go f.Run()
}

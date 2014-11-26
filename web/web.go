package web

import (
	"../future"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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
	d, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Write(d)
}

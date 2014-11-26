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

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, collection)
	case "POST":
		fmt.Fprintf(w, "POST")
	}
}

func Future(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	futureId := vars["future"]

	switch r.Method {
	case "GET":
		f := future.Get(futureId)
		d, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return
		}
		w.Write(d)
	}
}

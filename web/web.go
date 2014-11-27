package web

import (
	"../future"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

/**
 *
 */
func Collection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	collection := vars["collection"]

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, collection)
	case "POST":
		r.ParseForm()
		form := r.Form

		scheduled := form.Get("Scheduled")
		command := form.Get("Command")

		if command == "" {
			w.WriteHeader(405)
			return
		}

		t, err := time.Parse(time.RFC3339, scheduled)
		if err != nil {
			w.WriteHeader(405)
			return
		}

		f := future.New(t, command)
		f.Save()
		go f.Schedule() /* Fix me */

		d, _ := json.Marshal(f)
		w.Write(d)
	default:
		w.WriteHeader(405)
		return
	}
}

/**
 *
 */
func Future(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	futureId := vars["future"]

	switch r.Method {
	case "GET":
		f := future.Get(futureId)
		d, err := json.Marshal(f)
		if err != nil {
			return
		}
		w.Write(d)
	default:
		w.WriteHeader(405)
	}
}

/**
 *
 */
func FutureCancel(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// futureId := vars["future"]

	// switch r.Method {
	// default:
	// 	w.WriteHeader(405)
	// }
}

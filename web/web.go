package web

import (
	"../future"
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
	"time"
)

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

		t, err := time.Parse(time.RFC3339, scheduled)
		if err != nil {
			log.Fatal(err)
			return
		}

		f := future.Future{
			Id:        uuid.New(),
			Scheduled: t,
			Cancelled: false,
			Done:      false,
			Error:     true,
			Command:   *exec.Command("touch", "foo"),
		}

		d, _ := json.Marshal(f)
		w.Write(d)
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

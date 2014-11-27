package web

import (
	"github.com/gorilla/mux"
)

/**
 *
 */
func Routes() *mux.Router {
	rtr := mux.NewRouter()

	/* Collection routes */
	rtr.HandleFunc("/{collection}/", Collection).Methods("GET", "POST")

	/* Future routes */
	rtr.HandleFunc("/{collection}/{future}/", Future).Methods("GET")

	// rtr.HandleFunc("/{collection}/{future}/wait/", FutureWait).Methods("GET")
	rtr.HandleFunc("/{collection}/{future}/cancel/", FutureCancel).Methods("GET")
	// rtr.HandleFunc("/{collection}/{future}/run/", FutureRun).Methods("GET")

	return rtr
}

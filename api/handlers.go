package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/calculator/add", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":80", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!")
}

// curl -X POST -d '{"numbers": "4,6"}' http://localhost/calculator/add

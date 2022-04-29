package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"action-test/calculator"

	"github.com/gorilla/mux"
)

type input struct {
	Numbers string `json:"numbers"`
}

type output struct {
	Result  string
	Status  string
	Message string
}

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/calculator/add", add).Methods("POST")
	http.ListenAndServe(":80", r)
}

func add(w http.ResponseWriter, r *http.Request) {
	dataDecoder := json.NewDecoder(r.Body)
	numbers := input{}
	_ = dataDecoder.Decode(&numbers)

	calculator.Add(numbers.Numbers)
	//fmt.Println("adding numbers ", *sum)
	//fmt.Fprintf(w, "%v\n", *sum)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!\n")
}

// curl -X POST -d '{"numbers": "4,6"}' http://localhost/calculator/add

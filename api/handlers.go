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
	r.HandleFunc("/calculator/subtract", subtract).Methods("POST")
	r.HandleFunc("/calculator/multiply", multiply).Methods("POST")
	r.HandleFunc("/calculator/divide", divide).Methods("POST")
	http.ListenAndServe(":80", r)
}

func writeResponse(w http.ResponseWriter, result *float64, err error) {
	out := output{}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		out.Status = http.StatusText(http.StatusBadRequest)
		out.Message = err.Error()
	} else {
		w.WriteHeader(http.StatusOK)
		out.Result = fmt.Sprintf("%v", *result)
		out.Status = http.StatusText(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func add(w http.ResponseWriter, r *http.Request) {
	dataDecoder := json.NewDecoder(r.Body)
	numbers := input{}
	_ = dataDecoder.Decode(&numbers)

	sum, err := calculator.Add(numbers.Numbers)
	writeResponse(w, sum, err)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	dataDecoder := json.NewDecoder(r.Body)
	numbers := input{}
	_ = dataDecoder.Decode(&numbers)

	diff, err := calculator.Subtract(numbers.Numbers)
	writeResponse(w, diff, err)
}

func multiply(w http.ResponseWriter, r *http.Request) {
	dataDecoder := json.NewDecoder(r.Body)
	numbers := input{}
	_ = dataDecoder.Decode(&numbers)

	prod, err := calculator.Multiply(numbers.Numbers)
	writeResponse(w, prod, err)
}

func divide(w http.ResponseWriter, r *http.Request) {
	dataDecoder := json.NewDecoder(r.Body)
	numbers := input{}
	_ = dataDecoder.Decode(&numbers)

	sum, err := calculator.Divide(numbers.Numbers)
	writeResponse(w, sum, err)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!\n")
}

// curl -X POST -d '{"numbers": "4,6"}' http://localhost/calculator/add
// {"status": "success", "result": nil, "message": ""}

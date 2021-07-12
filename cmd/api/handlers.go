package main

import (
	"encoding/json"
	"erozen/mutants/internal/mutants"
	"io/ioutil"
	"net/http"
)

func responseError(w http.ResponseWriter, status int, code, message string) {
	type error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	b, err := json.Marshal(error{Code: code, Message: message})
	if err != nil {
		w.Write([]byte(message))
		return
	}

	w.Write(b)
}

func Mutant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusBadRequest, "invalid_body", err.Error())
		return
	}

	type request struct {
		DNA []string `json:"dna"`
	}

	var rq request
	if err := json.Unmarshal(body, &rq); err != nil {
		responseError(w, http.StatusBadRequest, "invalid_request", err.Error())
		return
	}

	if mutants.IsMutant(rq.DNA) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	return
}

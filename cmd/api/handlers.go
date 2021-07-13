package main

import (
	"encoding/json"
	"erozen/mutants/internal/mutants"
	"io/ioutil"
	"net/http"
)

type MutantHandler struct {
	service *mutants.Service
}

func (h *MutantHandler) Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.service.Stats(r.Context())
	if err != nil {
		responseError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	responseJSON(w, http.StatusOK, stats)
}

func (h *MutantHandler) Mutant(w http.ResponseWriter, r *http.Request) {
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

	isMutant, err := h.service.IsMutant(r.Context(), rq.DNA)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "internal_error", err.Error())
		return
	}

	if isMutant {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	return
}

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

func responseJSON(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	b, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(b)
}

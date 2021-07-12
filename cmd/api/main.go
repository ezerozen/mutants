package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Post("/mutant", Mutant)
	http.ListenAndServe(":8080", r)
}

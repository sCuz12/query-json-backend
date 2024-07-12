package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {
		// add any api routes here
		r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
			var payload struct {
				Content string `json:"content"`
			}
			payload.Content = "Hello, world"
			a.App.WriteJSON(w, http.StatusOK, payload)
		})
		r.Post("/query-json", func(w http.ResponseWriter, r *http.Request) {
			a.Handlers.JsonSearch(w,r)
		})
	})


	return r
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	// //Setup cors
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow your frontend origin
	    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	    ExposedHeaders:   []string{"Link"},
	    AllowCredentials: true,
	    MaxAge:           300, // Maximum value for Access-Control-Max-Age header in seconds
	})

	//use cors middleware
	r.Use(cors.Handler)

	r.Route("/api", func(mux chi.Router) {
		r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
			var payload struct {
				Content string `json:"content"`
			}
			payload.Content = "Hello, world"
			a.App.WriteJSON(w, http.StatusOK, payload)
		})

		r.Post("/query-json", func(w http.ResponseWriter, r *http.Request) {
			a.Handlers.JsonSearch(w, r)
		})

		r.Post("/query-recommendations", func(w http.ResponseWriter, r *http.Request) {
			a.Handlers.QueryRecommandations(w, r)
		})

		r.Post("/store-file", func(w http.ResponseWriter, r *http.Request) {
			a.Handlers.StoreFile(w, r)
		})

		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			a.Handlers.Register(w, r)
		})
	})

	return r
}

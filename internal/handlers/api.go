package handlers

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/hashcott/go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddleware.StripSlashes)

	r.Route("/account", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance)
	})
}

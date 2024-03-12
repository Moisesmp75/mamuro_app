package controllers

import (
	"mamuro_app/services"

	"github.com/go-chi/chi/v5"
)

func EmailController(r *chi.Mux) {
	r.Route("/mail", func(r chi.Router) {
		r.Post("/search", services.SearchData)
		r.Post("/", services.PostMail)
	})
}

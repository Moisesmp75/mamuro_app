package controllers

import "github.com/go-chi/chi/v5"

func AddControllers(r *chi.Mux) {
	EmailController(r)
}

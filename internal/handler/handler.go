package handler

import "github.com/go-chi/chi/v5"

func RegisterRouters(r *chi.Mux) {
	home := homeHandler{}
	r.Get("/", home.handleIndex)
}

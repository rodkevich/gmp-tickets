package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rodkevich/gmp-tickets/internal/ticket"
)

func RegisterRoutes(router *chi.Mux, validator *validator.Validate, usage ticket.UsageSchema) *Handler {
	h := NewHandler(usage, validator)

	router.Route("/api/v1/tickets", func(router chi.Router) {
		router.Get("/", h.List)
		router.Get("/{ticketID}", h.Get)
		router.Post("/", h.Create)
		router.Put("/{ticketID}", h.Update)
		router.Delete("/{ticketID}", h.Delete)
	})
	return h
}

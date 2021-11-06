package rest

import (
	"context"
	"database/sql"
	"github.com/rodkevich/gmp-tickets/lib/msg"
	"github.com/rodkevich/gmp-tickets/lib/validation"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rodkevich/gmp-tickets/internal/ticket"
)

type HTTP interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	validation *validator.Validate
	usage      ticket.UsageSchema
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req ticket.CreationRequest
	err := ticket.Bind(r.Body, &req)
	if err != nil {
		msg.ReturnClientError(w, err.Error())
	}
	errs := validation.Validate(h.validation, req)
	if errs != nil {
		msg.ReturnClientError(w, err.Error())
		return
	}
	t := ticket.Parse(&req)

	tk, err := h.usage.Create(context.Background(), t)
	if err != nil {
		if err == sql.ErrNoRows {
			msg.ReturnClientError(w, err.Error())
			return
		}
		msg.ReturnServerError(w, err)
		return
	}
	log.Println(tk)
	response := ticket.RtnPrepare(*tk, true, nil)
	if err != nil {
		msg.ReturnServerError(w, err)
		return
	}
	msg.ReturnJSON(w, response)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewHandler(usage ticket.UsageSchema, validation *validator.Validate) *Handler {
	return &Handler{
		usage:      usage,
		validation: validation,
	}
}

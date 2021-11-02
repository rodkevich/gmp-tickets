package rest

import (
	"encoding/json"
	"golang.org/x/net/context"
	"io/ioutil"
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

func NewHandler(usage ticket.UsageSchema, validation *validator.Validate) *Handler {
	return &Handler{
		usage:      usage,
		validation: validation,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req ticket.CreationRequest
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(b, &req); err != nil {
		log.Println(err)
		return
	}

	resp, err := h.usage.Create(context.Background(), req)
	if err != nil {
		return
	}
	log.Println(resp)
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

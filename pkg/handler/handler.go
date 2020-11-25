package handler

import (
	"github.com/gorilla/mux"
	"github.com/sashindionicus/DBDocument/pkg/service"
)

type Handler struct {
	documentsService service.Documents
}

func NewHandler(documentsService service.Documents) *Handler {
	return &Handler{
		documentsService: documentsService,
	}
}

func (h *Handler) Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/documents/{id}", h.getDocument).Methods("GET")
	r.HandleFunc("/documents", h.createDocument).Methods("POST")
	r.HandleFunc("/documents/{id}", h.deleteDocument).Methods("DELETE")

	return r
}

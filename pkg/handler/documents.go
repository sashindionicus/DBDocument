package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sashindionicus/DBDocument"
	"net/http"
	"strconv"
)

type Error struct {
	Message string `json:"message"`
}

func (h *Handler) getDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	doc, err := h.documentsService.Get(id)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&doc)
}

func (h *Handler) createDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var document DBDocument.Document
	err := json.NewDecoder(r.Body).Decode(&document)
	if err != nil {
		badRequest(w, err)
		return
	}

	id, err := h.documentsService.Create(document)
	if err != nil {
		badRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		badRequest(w, err)
		return
	}

	err = h.documentsService.Delete(id)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func badRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&Error{Message: err.Error()})
}

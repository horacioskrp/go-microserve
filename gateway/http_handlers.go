package main

import "net/http"

type handler struct {
}

func NewHander() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/custermers/{customerID}/orders", h.HandleCreatorOrder)
}

func (h *handler) HandleCreatorOrder(w http.ResponseWriter, r *http.Request) {

}

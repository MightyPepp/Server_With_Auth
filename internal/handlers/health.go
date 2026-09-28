package handlers

import "net/http"

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func RegisterHealthHandler(mux *http.ServeMux, hh *HealthHandler) {
	mux.HandleFunc("/health", hh.pingServer)
}

func (h *HealthHandler) pingServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running!"))
}
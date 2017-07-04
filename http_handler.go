package main

import (
	"encoding/json"
	"net/http"
)

// HTTPHandler represents an HTTP handler for process listing
type HTTPHandler struct {
	processLister ProcessLister
}

// ServeHTTP handles the HTTP request/response process
func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.processLister.List())
}

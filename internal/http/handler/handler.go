package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
	AppName string
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprintf(w, "Hello to, %s!", h.AppName)
}

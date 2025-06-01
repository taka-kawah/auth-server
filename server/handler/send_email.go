package handler

import (
	"fmt"
	"net/http"
)

func (h *HttpHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"status":"permits only GET"}`, http.StatusMethodNotAllowed)
		return
	}
	mailAddress := r.URL.Query().Get("email")
	id, err := h.r.Create(mailAddress)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":%s}`, err), http.StatusInternalServerError)
		return
	}
	if err := h.a.SendEmail(mailAddress, id); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":%s}`, err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

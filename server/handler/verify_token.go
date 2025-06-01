package handler

import (
	"fmt"
	"net/http"
)

func (h *HttpHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"status":"permits only POST"}`, http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authorization")
	if _, err := h.tk.ExtractId(token); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

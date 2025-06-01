package handler

import (
	"auth-server/util"
	"fmt"
	"net/http"
)

func (h *HttpHandler) UpdateMailAddress(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, util.TokenNotSetMessage), http.StatusForbidden)
		return
	}
	_, err := h.tk.ExtractId(token)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusForbidden)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, `{"status":"permits only GET"}`, http.StatusMethodNotAllowed)
		return
	}
	email := r.URL.Query().Get("mail")
	newEmail := r.URL.Query().Get("new_email")
	if err = h.a.UpdateMailAddress(email, newEmail); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

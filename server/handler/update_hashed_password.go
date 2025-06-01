package handler

import (
	"auth-server/util"
	"fmt"
	"net/http"
)

func (h *HttpHandler) UpdateHashedPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"status":"permits only GET"}`, http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, util.TokenNotSetMessage), http.StatusForbidden)
		return
	}
	id, err := h.tk.ExtractId(token)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusForbidden)
		return
	}
	hashedPassword := r.URL.Query().Get("hashed_password")
	if err = h.a.UpdateHashedPassword(id, hashedPassword); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

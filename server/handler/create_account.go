package handler

import (
	"auth-server/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *HttpHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"status":"permits only POST"}`, http.StatusMethodNotAllowed)
		return
	}
	var a models.Account
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	if err := h.v.Validate(a); err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusForbidden)
		return
	}
	id, err := h.a.Create(a.MailAddress, a.HashedPassword)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	token, err := h.tk.ProvideJWT(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Authorization": token})
}

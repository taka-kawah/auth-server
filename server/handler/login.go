package handler

import (
	"auth-server/models"
	"auth-server/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *HttpHandler) Login(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusBadRequest)
		return
	}
	id, err := h.a.Login(a.MailAddress, a.HashedPassword)
	if err != nil && err.Error() == util.NotFoundMessage {
		http.Error(w, fmt.Sprintf(`{"status":"%s"}`, util.NotFoundMessage), http.StatusNotFound)
		return
	}
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

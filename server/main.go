package main

import (
	"auth-server/handler"
	accesstoken "auth-server/interfaces/access_token"
	googleoauth "auth-server/interfaces/google_oauth"
	mailinterface "auth-server/interfaces/mail_interface"
	"auth-server/interfaces/repository"
	"auth-server/services/tables"
	"auth-server/services/tokens"
	"auth-server/services/validation"
	"context"
	"log"
	"net/http"
)

func main() {
	sqlDb, err := repository.CreatePostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()
	acRepo := repository.NewAccountRepository(context.Background(), sqlDb)
	reRepo := repository.NewReserveRepository(context.Background(), sqlDb)
	mailSender := mailinterface.NewMailSender()
	acService := tables.NewAccountService(mailSender, acRepo)
	reService := tables.NewReserveService(reRepo)
	v := validation.New()
	t := accesstoken.NewTokenManager()
	g := googleoauth.NewGoogleTokenVerifier(context.Background())
	tk := tokens.NewTokenManager(t, g)
	h := handler.NewHttpHandler(v, acService, reService, tk)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("古舘くれあさんに、フル勃ち...w"))
	})
	mux.HandleFunc("/signup", h.CreateAccount)
	mux.HandleFunc("/signin", h.Login)
	mux.HandleFunc("/mail", h.SendEmail)
	mux.HandleFunc("/email_change", h.UpdateMailAddress)
	mux.HandleFunc("/password_change", h.UpdateHashedPassword)
	mux.HandleFunc("token_verification/", h.VerifyToken)

	http.Handle("/auth/", http.StripPrefix("/auth", mux))

	http.ListenAndServe(":8000", mux)
}

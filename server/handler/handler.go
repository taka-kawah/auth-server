package handler

import (
	"auth-server/services/tables"
	"auth-server/services/tokens"
	"auth-server/services/validation"
)

type HttpHandler struct {
	v  validation.ValidationService
	a  tables.AccountService
	r  tables.ReserveService
	tk tokens.TokenManager
}

func NewHttpHandler(v validation.ValidationService, a tables.AccountService, r tables.ReserveService, tk tokens.TokenManager) *HttpHandler {
	return &HttpHandler{v: v, a: a, r: r, tk: tk}
}

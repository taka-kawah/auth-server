package googleoauth

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"
)

type GoogleTokenVerifier interface {
	ExtractId(token string) (string, error)
}

type googleTokenVerifierImpl struct {
	ctx context.Context
}

func NewGoogleTokenVerifier(ctx context.Context) *googleTokenVerifierImpl {
	return &googleTokenVerifierImpl{ctx: ctx}
}

func (g *googleTokenVerifierImpl) ExtractId(token string) (string, error) {
	payload, err := idtoken.Validate(g.ctx, token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return "", nil
	}
	sub := payload.Claims["sub"]
	id := sub.(string)
	return id, nil
}

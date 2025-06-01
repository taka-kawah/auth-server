package tokens

import (
	accesstoken "auth-server/interfaces/access_token"
	googleoauth "auth-server/interfaces/google_oauth"
)

type TokenManager interface {
	ProvideJWT(id string) (string, error)
	ExtractId(token string) (string, error)
}

type tokenManagerImpl struct {
	j accesstoken.TokenManager
	g googleoauth.GoogleTokenVerifier
}

func NewTokenManager(j accesstoken.TokenManager, g googleoauth.GoogleTokenVerifier) *tokenManagerImpl {
	return &tokenManagerImpl{j: j, g: g}
}

func (t *tokenManagerImpl) ProvideJWT(id string) (string, error) {
	return t.j.Provide(id)
}

func (t *tokenManagerImpl) ExtractId(token string) (string, error) {
	var (
		id  string
		err error
	)
	switch len(token) {
	case 26:
		id, err = t.j.ExtractId(token)
	case 27:
		id, err = t.g.ExtractId(token)
	}
	if err != nil {
		return "", err
	}
	return id, nil
}

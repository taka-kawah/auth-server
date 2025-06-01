package accesstoken

import (
	"auth-server/util"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	Provide(mailAddress string) (string, error)
	ExtractId(tokenString string) (string, error)
}

type tokenManagerImpl struct{}

func NewTokenManager() *tokenManagerImpl {
	return &tokenManagerImpl{}
}

func (tm *tokenManagerImpl) Provide(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (tm *tokenManagerImpl) ExtractId(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", errors.New(util.TokenExpiredMessage)
		}
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return string(claims["user_id"].(string)), nil
	}
	return "", errors.New("failed to get token claims")
}

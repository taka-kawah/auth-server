package util

import "fmt"

const (
	NotFoundMessage     string = "not found"
	TokenExpiredMessage string = "token expired"
	TokenNotSetMessage  string = "token not set"
)

func RequiredMessage(v string) string {
	return fmt.Sprintf("%s is required", v)
}

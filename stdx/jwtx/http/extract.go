package http

import (
	"errors"
	"net/http"

	"github.com/evgenivanovi/gpl/stdx/jwtx"
	"github.com/golang-jwt/jwt/v5/request"
)

// CookieExtractor extracts the JWT token from the cookie using the passed in cookie name.
type CookieExtractor string

func (e CookieExtractor) String() string {
	return string(e)
}

func (e CookieExtractor) ExtractToken(req *http.Request) (string, error) {

	cookie, err := req.Cookie(e.String())

	if cookie == nil {
		return "", request.ErrNoTokenInRequest
	}

	if cookie.Value == "" {
		return "", request.ErrNoTokenInRequest
	}

	if errors.Is(err, http.ErrNoCookie) {
		return "", request.ErrNoTokenInRequest
	}

	return cookie.Value, nil

}

// ContextExtractor extracts the JWT token from the request's context.Context.
type ContextExtractor struct{}

func (e ContextExtractor) ExtractToken(req *http.Request) (string, error) {
	token := jwtx.FromCtxAsString(req.Context())
	if token == "" {
		return "", request.ErrNoTokenInRequest
	}
	return token, nil
}

type GenerateExtractor struct {
	generate func() (string, error)
}

func NewGenerateExtractor(
	generate func() (string, error),
) GenerateExtractor {
	return GenerateExtractor{
		generate: generate,
	}
}

func NewGenerateExtractorWithValue(token string) GenerateExtractor {
	return NewGenerateExtractor(
		func() (string, error) {
			return token, nil
		},
	)
}

func (e GenerateExtractor) ExtractToken(_ *http.Request) (string, error) {
	return e.generate()
}

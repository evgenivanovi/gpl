package jwtx

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5/request"
)

/* __________________________________________________ */

// CookieExtractor extracts the JWT token from the cookie using the passed in cookieName.
type CookieExtractor string

func (e CookieExtractor) String() string {
	return string(e)
}

func (e CookieExtractor) ExtractToken(req *http.Request) (string, error) {
	cookie, err := req.Cookie(e.String())
	if cookie != nil && cookie.Value != "" && !errors.Is(err, http.ErrNoCookie) {
		return cookie.Value, nil
	}
	return "", request.ErrNoTokenInRequest
}

/* __________________________________________________ */

// ContextExtractor extracts the JWT token from the request's context.Context.
type ContextExtractor struct{}

func (e ContextExtractor) ExtractToken(req *http.Request) (string, error) {
	token := FromCtxAsString(req.Context())
	if token == "" {
		return "", request.ErrNoTokenInRequest
	}
	return token, nil
}

/* __________________________________________________ */

package http

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

type MiddlewareOp func(*middleware)

func WithKey(key jwt.Keyfunc) MiddlewareOp {
	return func(mw *middleware) {
		mw.key = key
	}
}

func WithMethod(method jwt.SigningMethod) MiddlewareOp {
	return func(mw *middleware) {
		mw.method = method
	}
}

func WithClaims(claims func() jwt.Claims) MiddlewareOp {
	return func(mw *middleware) {
		mw.claims = claims
	}
}

func WithExtractor(extractors ...request.Extractor) MiddlewareOp {
	return func(mw *middleware) {
		mw.extractor = extractors
	}
}

func WithGenerator(generator func() string) MiddlewareOp {
	return func(mw *middleware) {
		mw.generator = generator
	}
}

func WithVerifier(verifier func(http.ResponseWriter, *http.Request, *jwt.Token, string) error) MiddlewareOp {
	return func(mw *middleware) {
		mw.verifier = verifier
	}
}

func WithRecoverer(recoverer func(http.ResponseWriter, *http.Request, error)) MiddlewareOp {
	return func(mw *middleware) {
		mw.recoverer = recoverer
	}
}

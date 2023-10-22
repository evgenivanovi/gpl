package jwtx

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

/* __________________________________________________ */

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
		mw.extractors = extractors
	}
}

func WithBefore(before func(http.ResponseWriter, *http.Request)) MiddlewareOp {
	return func(mw *middleware) {
		mw.before = before
	}
}

func WithAfter(after func(http.ResponseWriter, *http.Request, *jwt.Token, string) error) MiddlewareOp {
	return func(mw *middleware) {
		mw.after = after
	}
}

func WithRecover(recover func(http.ResponseWriter, *http.Request, error)) MiddlewareOp {
	return func(mw *middleware) {
		mw.recover = recover
	}
}

/* __________________________________________________ */

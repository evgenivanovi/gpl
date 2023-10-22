package grpc

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
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

func WithExtractor(extractors ...Extractor) MiddlewareOp {
	return func(mw *middleware) {
		mw.extractors = extractors
	}
}

func WithGenerator(generator func() string) MiddlewareOp {
	return func(mw *middleware) {
		mw.generator = generator
	}
}

func WithVerifier(verifier func(context.Context, *jwt.Token, string) (context.Context, error)) MiddlewareOp {
	return func(mw *middleware) {
		mw.verifier = verifier
	}
}

func WithRecoverer(recoverer UnaryServerErrorInterceptor) MiddlewareOp {
	return func(mw *middleware) {
		mw.recoverer = recoverer
	}
}

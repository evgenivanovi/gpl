package jwtx

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const jwtTokenCtxKey contextKey = "ctx.jwt.token"
const jwtTokenStringCtxKey contextKey = "ctx.jwt.token.string"

// FromCtx
// Takes a context.Context and returns the JWT token associated with it (if any).
func FromCtx(ctx context.Context) *jwt.Token {
	if value, ok := ctx.Value(jwtTokenCtxKey).(*jwt.Token); ok {
		return value
	}
	return nil
}

// FromCtxAsString
// Takes a context.Context and returns the JWT token associated with it (if any).
func FromCtxAsString(ctx context.Context) string {
	if value, ok := ctx.Value(jwtTokenStringCtxKey).(string); ok {
		return value
	}
	return ""
}

// WithCtx
// Associates a JWT token with a context.Context and returns it.
func WithCtx(ctx context.Context, token *jwt.Token) context.Context {
	value := FromCtx(ctx)
	if value == token {
		return ctx
	}
	return context.WithValue(ctx, jwtTokenCtxKey, token)
}

func WithRequestCtx(request *http.Request, token *jwt.Token) {
	ctx := WithCtx(request.Context(), token)
	*request = *request.WithContext(ctx)
}

// WithCtxAsString
// Associates a JWT token with a context.Context and returns it.
func WithCtxAsString(ctx context.Context, token string) context.Context {
	value := FromCtxAsString(ctx)
	if value == token {
		return ctx
	}
	return context.WithValue(ctx, jwtTokenStringCtxKey, token)
}

func WithRequestCtxAsString(request *http.Request, token string) {
	ctx := WithCtxAsString(request.Context(), token)
	*request = *request.WithContext(ctx)
}

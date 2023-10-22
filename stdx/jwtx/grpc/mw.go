package grpc

import (
	"context"
	"errors"

	"github.com/evgenivanovi/gpl/stdx/jwtx"
	"github.com/golang-jwt/jwt/v5"
	jwtreq "github.com/golang-jwt/jwt/v5/request"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type middleware struct {
	key    jwt.Keyfunc
	method jwt.SigningMethod
	claims func() jwt.Claims

	extractors MultiExtractor
	generator  func() string

	verifier  func(context.Context, *jwt.Token, string) (context.Context, error)
	recoverer UnaryServerErrorInterceptor
}

func NewUnary(ops ...MiddlewareOp) grpc.UnaryServerInterceptor {

	obj := defaultMiddleware()

	for _, op := range ops {
		op(obj)
	}

	if obj.method == nil {
		panic("signing method is required")
	}

	if obj.key == nil {
		panic("key function is required")
	}

	if len(obj.extractors) == 0 {
		panic("extractors should not be empty")
	}

	return obj.wrap()

}

func defaultMiddleware() *middleware {
	return &middleware{
		claims: func() jwt.Claims {
			return make(jwt.MapClaims)
		},
		recoverer: func(
			ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, ex error,
		) (resp any, err error) {
			return nil, status.Error(codes.Unauthenticated, "")
		},
	}
}

func (mw *middleware) wrap() grpc.UnaryServerInterceptor {

	parser := jwt.NewParser(
		jwt.WithValidMethods(
			[]string{
				mw.method.Alg(),
			},
		),
	)

	isErrorNoToken := func(err error) bool {
		return errors.Is(err, jwtreq.ErrNoTokenInRequest)
	}

	return func(
		ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp any, err error) {

		tokenString, err := mw.extractors.ExtractToken(ctx)

		if err != nil && isErrorNoToken(err) && mw.generator == nil {
			return mw.recoverer(ctx, req, info, handler, err)
		}

		if err != nil && isErrorNoToken(err) && mw.generator != nil {
			tokenString = mw.generator()
		}

		token, err := parser.ParseWithClaims(tokenString, mw.claims(), mw.key)
		if err != nil {
			return mw.recoverer(ctx, req, info, handler, err)
		}

		ctx = withJWTContext(ctx, token, tokenString)

		ctx, err = mw.verifier(ctx, token, tokenString)
		if err != nil {
			return mw.recoverer(ctx, req, info, handler, err)
		}

		return handler(ctx, req)

	}

}

func withJWTContext(ctx context.Context, token *jwt.Token, tokenString string) context.Context {
	if token != nil {
		ctx = jwtx.WithCtx(ctx, token)
	}
	if tokenString != "" {
		ctx = jwtx.WithCtxAsString(ctx, tokenString)
	}
	return ctx
}

package jwtx

import (
	"net/http"

	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

/* __________________________________________________ */

type middleware struct {
	key        jwt.Keyfunc
	method     jwt.SigningMethod
	claims     func() jwt.Claims
	extractors request.MultiExtractor

	before  func(http.ResponseWriter, *http.Request)
	after   func(http.ResponseWriter, *http.Request, *jwt.Token, string) error
	recover func(http.ResponseWriter, *http.Request, error)
}

func New(opts ...MiddlewareOp) func(next http.Handler) http.Handler {

	mw := defaultMiddleware()

	for _, opt := range opts {
		opt(mw)
	}

	if mw.method == nil {
		panic("signing method is required")
	}

	if mw.key == nil {
		panic("key function is required")
	}

	return mw.wrap

}

func defaultMiddleware() *middleware {
	return &middleware{
		claims: func() jwt.Claims {
			return make(jwt.MapClaims)
		},
		extractors: request.MultiExtractor{
			request.HeaderExtractor{
				headers.AuthorizationKey.String(),
			},
		},
		before: func(_ http.ResponseWriter, _ *http.Request) {},
		after:  func(_ http.ResponseWriter, _ *http.Request, _ *jwt.Token, _ string) error { return nil },
		recover: func(w http.ResponseWriter, _ *http.Request, _ error) {
			w.WriteHeader(http.StatusUnauthorized)
		},
	}
}

func (mw *middleware) wrap(next http.Handler) http.Handler {

	parser := jwt.NewParser(
		jwt.WithValidMethods(
			[]string{
				mw.method.Alg(),
			},
		),
	)

	mwFunc := func(writer http.ResponseWriter, request *http.Request) {
		mw.before(writer, request)

		tokenString, err := mw.extractors.ExtractToken(request)
		if err != nil {
			mw.recover(writer, request, err)
			return
		}

		token, err := parser.ParseWithClaims(tokenString, mw.claims(), mw.key)
		if err != nil {
			mw.recover(writer, request, err)
			return
		}

		withRequestCtx(request, token, tokenString)

		err = mw.after(writer, request, token, tokenString)
		if err != nil {
			mw.recover(writer, request, err)
			return
		}

		next.ServeHTTP(writer, request)
	}

	return http.HandlerFunc(mwFunc)

}

func withRequestCtx(request *http.Request, token *jwt.Token, tokenString string) {
	ctx := request.Context()
	if token != nil {
		ctx = WithCtx(ctx, token)
	}
	if tokenString != "" {
		ctx = WithCtxAsString(ctx, tokenString)
	}
	*request = *request.WithContext(ctx)
}

/* __________________________________________________ */

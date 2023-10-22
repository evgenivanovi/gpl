package http

import (
	"net/http"

	"github.com/evgenivanovi/gpl/stdx/jwtx"
	"github.com/evgenivanovi/gpl/stdx/mw"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/golang-jwt/jwt/v5"
	jwtreq "github.com/golang-jwt/jwt/v5/request"
	"github.com/pkg/errors"
)

type middleware struct {
	key    jwt.Keyfunc
	method jwt.SigningMethod
	claims func() jwt.Claims

	extractor jwtreq.MultiExtractor
	generator func() string

	verifier  func(http.ResponseWriter, *http.Request, *jwt.Token, string) error
	recoverer func(http.ResponseWriter, *http.Request, error)
}

func New(ops ...MiddlewareOp) mw.MW {

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

	return obj.wrap

}

func defaultMiddleware() *middleware {
	return &middleware{
		claims: func() jwt.Claims {
			return make(jwt.MapClaims)
		},
		extractor: jwtreq.MultiExtractor{
			jwtreq.HeaderExtractor{
				headers.AuthorizationKey.String(),
			},
		},
		verifier: func(_ http.ResponseWriter, _ *http.Request, _ *jwt.Token, _ string) error { return nil },
		recoverer: func(w http.ResponseWriter, _ *http.Request, _ error) {
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

	isErrorNoToken := func(err error) bool {
		return errors.Is(err, jwtreq.ErrNoTokenInRequest)
	}

	mwFunc := func(writer http.ResponseWriter, request *http.Request) {

		tokenString, err := mw.extractor.ExtractToken(request)

		if err != nil && isErrorNoToken(err) && mw.generator == nil {
			mw.recoverer(writer, request, err)
			return
		}

		if err != nil && isErrorNoToken(err) && mw.generator != nil {
			tokenString = mw.generator()
		}

		token, err := parser.ParseWithClaims(tokenString, mw.claims(), mw.key)
		if err != nil {
			mw.recoverer(writer, request, err)
			return
		}

		withJWTContext(request, token, tokenString)

		err = mw.verifier(writer, request, token, tokenString)
		if err != nil {
			mw.recoverer(writer, request, err)
			return
		}

		next.ServeHTTP(writer, request)

	}

	return http.HandlerFunc(mwFunc)

}

func withJWTContext(request *http.Request, token *jwt.Token, tokenString string) {
	ctx := request.Context()
	if token != nil {
		ctx = jwtx.WithCtx(ctx, token)
	}
	if tokenString != "" {
		ctx = jwtx.WithCtxAsString(ctx, tokenString)
	}
	*request = *request.WithContext(ctx)
}

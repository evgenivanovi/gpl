package mw

import (
	"net/http"
)

type (
	MW     func(next http.Handler) http.Handler
	MWFunc func(next http.HandlerFunc) http.HandlerFunc
)

func Conveyor(handler http.Handler, mws ...MW) http.Handler {
	for _, mw := range mws {
		handler = mw(handler)
	}
	return handler
}

func ConveyorFunc(handler http.HandlerFunc, mws ...MWFunc) http.HandlerFunc {
	for _, mw := range mws {
		handler = mw(handler)
	}
	return handler
}

package mw

import (
	"net/http"
	"time"

	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
)

const CallDurationKey = "http.call.duration"

const RequestURIKey = "http.request.uri"
const RequestMethodKey = "http.request.method"

const ResponseStatusKey = "http.response.status"
const ResponseSizeKey = "http.response.size"

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	// записываем ответ, используя оригинальный http.ResponseWriter
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	// записываем код статуса, используя оригинальный http.ResponseWriter
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}

func WithLogging(handler http.Handler) http.Handler {

	logFn := func(writer http.ResponseWriter, request *http.Request) {

		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lrw := loggingResponseWriter{
			ResponseWriter: writer,
			responseData:   responseData,
		}

		handler.ServeHTTP(&lrw, request)

		duration := time.Since(start)

		slogx.FromCtx(request.Context()).Debug(
			"Processed HTTP request.",
			RequestURIKey, request.RequestURI,
			RequestMethodKey, request.Method,
			ResponseStatusKey, responseData.status,
			ResponseSizeKey, responseData.size,
			CallDurationKey, duration,
		)

	}

	return http.HandlerFunc(logFn)

}

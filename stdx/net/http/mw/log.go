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

type loggingData struct {
	status int
	size   int
}

type loggingWriter struct {
	http.ResponseWriter
	loggingData *loggingData
}

func (r *loggingWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.loggingData.size += size
	return size, err
}

func (r *loggingWriter) WriteHeader(code int) {
	r.ResponseWriter.WriteHeader(code)
	r.loggingData.status = code
}

func WithLogging(handler http.Handler) http.Handler {

	log := func(writer http.ResponseWriter, request *http.Request) {

		start := time.Now()

		data := &loggingData{
			status: 0,
			size:   0,
		}

		lrw := loggingWriter{
			ResponseWriter: writer,
			loggingData:    data,
		}

		handler.ServeHTTP(&lrw, request)

		duration := time.Since(start)

		slogx.FromCtx(request.Context()).Debug(
			"Processed HTTP request.",
			RequestURIKey, request.RequestURI,
			RequestMethodKey, request.Method,
			ResponseStatusKey, data.status,
			ResponseSizeKey, data.size,
			CallDurationKey, duration,
		)

	}

	return http.HandlerFunc(log)

}

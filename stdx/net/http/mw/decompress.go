package mw

import (
	"compress/gzip"
	"io"
	"net/http"

	iox "github.com/evgenivanovi/gpl/std/io"
	"github.com/evgenivanovi/gpl/std/str"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/gookit/goutil"
)

// The following headers will be dropped from the request if decompression's applies.
// Their values will be moved to the corresponding X-Original- headers.
var requestHeadersToDrop = []string{
	headers.ContentEncodingKey.String(),
	headers.ContentLengthKey.String(),
}

func WithDecompress(next http.Handler) http.Handler {

	decompressFn := func(writer http.ResponseWriter, request *http.Request) {

		var readerCloser io.ReadCloser

		switch request.Header.Get(headers.ContentEncodingKey.String()) {
		case headers.EncodingGZIP.String():
			reader, err := gzip.NewReader(request.Body)
			if err == nil {
				readerCloser = reader
			} else {
				errorReader := iox.NewOnErrorReader(err)
				readerCloser = io.NopCloser(errorReader)
			}
		}

		if !goutil.IsNil(readerCloser) {
			request.Body = readerCloser
			for _, drop := range requestHeadersToDrop {
				if value := request.Header.Get(drop); value != "" {
					request.Header.Add(str.Join("X-Original-", drop), value)
					request.Header.Del(drop)
				}
			}
		}

		next.ServeHTTP(writer, request)

	}

	return http.HandlerFunc(decompressFn)

}

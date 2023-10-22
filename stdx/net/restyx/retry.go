package restyx

import (
	"net/http"
	"time"

	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var ErrNoRetry = errors.New("no retry")

func RetryOnManyRequestsAfterSeconds() resty.RetryAfterFunc {
	return func(client *resty.Client, response *resty.Response) (time.Duration, error) {

		retry := StatusFilter(http.StatusTooManyRequests)(response) &&
			HeaderFilter(headers.RetryAfterKey.String())(response)

		if !retry {
			return time.Duration(0), ErrNoRetry
		}

		delay, err := conv.MapInt64(RetryAfterValue(response))

		if err != nil {
			return time.Duration(0), ErrNoRetry
		}

		return time.Duration(delay) * time.Second, nil

	}
}

func RetryOnManyRequestsAfterDuration() resty.RetryAfterFunc {
	return func(client *resty.Client, response *resty.Response) (time.Duration, error) {

		retry := StatusFilter(http.StatusTooManyRequests)(response) &&
			HeaderFilter(headers.RetryAfterKey.String())(response)

		if !retry {
			return time.Duration(0), ErrNoRetry
		}

		return time.ParseDuration(RetryAfterValue(response))

	}
}

func RetryAfterValue(response *resty.Response) string {
	return response.
		Header().
		Get(headers.RetryAfterKey.String())
}

package restyx

import (
	"net/http"
	"time"

	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

func RetryOnManyRequestsAfterSeconds() resty.RetryAfterFunc {
	return func(client *resty.Client, response *resty.Response) (time.Duration, error) {

		retry :=
			StatusFilter(http.StatusTooManyRequests)(response) &&
				HeaderFilter(headers.RetryAfterKey.String())(response)

		if !retry {
			return defaultRetryResponse()
		}

		delay, delayErr := conv.MapInt64(RetryAfterValue(response))
		if delayErr != nil {
			return defaultRetryResponse()
		}

		var timeout = time.Duration(delay)
		return timeout * time.Second, nil

	}
}

func RetryOnManyRequestsAfterDuration() resty.RetryAfterFunc {
	return func(client *resty.Client, response *resty.Response) (time.Duration, error) {

		retry :=
			StatusFilter(http.StatusTooManyRequests)(response) &&
				HeaderFilter(headers.RetryAfterKey.String())(response)

		if !retry {
			return defaultRetryResponse()
		}

		return time.ParseDuration(RetryAfterValue(response))

	}
}

func defaultRetryResponse() (time.Duration, error) {
	var timeout time.Duration
	return timeout, errors.New("no retry")
}

/* __________________________________________________ */

func RetryAfterValue(response *resty.Response) string {
	return response.
		Header().
		Get(headers.RetryAfterKey.String())
}

/* __________________________________________________ */

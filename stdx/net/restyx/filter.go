package restyx

import "github.com/go-resty/resty/v2"

/* __________________________________________________ */

func StatusFilter(status int) func(response *resty.Response) bool {
	return func(response *resty.Response) bool {
		return response.StatusCode() == status
	}
}

func HeaderFilter(header string) func(response *resty.Response) bool {
	return func(response *resty.Response) bool {
		return response.Header().Get(header) != ""
	}
}

/* __________________________________________________ */

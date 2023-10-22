package restyx

import (
	"github.com/evgenivanovi/gpl/std"
	"github.com/go-resty/resty/v2"
)

func StatusFilter(status int) func(response *resty.Response) bool {
	return func(response *resty.Response) bool {
		return response.StatusCode() == status
	}
}

func HeaderFilter(header string) func(response *resty.Response) bool {
	return func(response *resty.Response) bool {
		return response.Header().Get(header) != std.Empty
	}
}

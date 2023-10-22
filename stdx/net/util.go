package net

import (
	"net/url"

	"github.com/evgenivanovi/gpl/std"
)

func RawURL(value string) string {
	u, err := url.Parse(value)
	if err != nil {
		return value
	}

	u.RawQuery = std.Empty
	return u.String()
}

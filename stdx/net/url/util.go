package url

import (
	"net/url"
)

func Raw(value string) string {
	u, err := url.Parse(value)

	if err != nil {
		return value
	}

	u.RawQuery = ""
	return u.String()
}

package jwtx

import (
	"strings"
)

/* __________________________________________________ */

func ExtractBearerToken(token string) (string, error) {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:], nil
	}
	return token, nil
}

/* __________________________________________________ */

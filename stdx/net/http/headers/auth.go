package headers

import (
	"fmt"
	"strings"

	"github.com/evgenivanovi/gpl/std"
	"github.com/evgenivanovi/gpl/std/str"
)

/* __________________________________________________ */

const (
	AuthorizationKey HeaderKey = "Authorization"

	TokenTypeUnknown TokenType = "unknown"
	TokenTypeBearer  TokenType = "bearer"
	TokenTypeMAC     TokenType = "mac"
)

/* __________________________________________________ */

type TokenType HeaderValue

func (tt TokenType) String() string {
	return string(tt)
}

func AuthorizationTokenType(token string) TokenType {

	isBearer := func(token string) bool {
		return strings.ToLower(token[:len(TokenTypeBearer)]) == TokenTypeBearer.String()
	}

	if isBearer(token) {
		return TokenTypeBearer
	}

	isMac := func(token string) bool {
		return strings.ToLower(token[:len(TokenTypeMAC)]) == TokenTypeMAC.String()
	}

	if isMac(token) {
		return TokenTypeMAC
	}

	return TokenTypeUnknown

}

/* __________________________________________________ */

func BuildBearerToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

func ExtractBearerToken(token string) (string, error) {
	if len(token) > 6 && strings.ToLower(token[0:7]) == str.Join(TokenTypeBearer.String(), std.Space) {
		return token[7:], nil
	}
	return token, nil
}

/* __________________________________________________ */

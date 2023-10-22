package jwtx

import "github.com/golang-jwt/jwt/v5"

/* __________________________________________________ */

const ClaimJWTKey = "jti"
const ClaimIssuerKey = "iss"
const ClaimSubjectKey = "sub"
const ClaimAudienceKey = "aud"
const ClaimExpiresAtKey = "exp"
const ClaimNotBeforeKey = "nbf"
const ClaimIssuedAtKey = "iat"

/* __________________________________________________ */

func SignJWT(token jwt.Token, secret string) (string, error) {
	return token.SignedString([]byte(secret))
}

/* __________________________________________________ */

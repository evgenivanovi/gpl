package http

import (
	"github.com/evgenivanovi/gpl/stdx/jwtx"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/golang-jwt/jwt/v5/request"
)

// AuthorizationHeaderExtractor extracts a token from Authorization header
var AuthorizationHeaderExtractor = &request.HeaderExtractor{
	headers.AuthorizationKey.String(),
}

// BearerAuthorizationHeaderExtractor extracts a bearer token from Authorization header
// Uses PostExtractionFilter to strip "Bearer " prefix from header
var BearerAuthorizationHeaderExtractor = &request.PostExtractionFilter{
	Extractor: AuthorizationHeaderExtractor,
	Filter:    jwtx.ExtractBearerToken,
}

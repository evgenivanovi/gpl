package grpc

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5/request"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

// Extractor is an interface for extracting a token from request.
// The ExtractToken method should return a token string or an error.
// If no token is present, you must return ErrNoTokenInRequest.
type Extractor interface {
	ExtractToken(ctx context.Context) (string, error)
}

// MultiExtractor tries Extractors in order until one returns a token string or an error occurs
type MultiExtractor []Extractor

func (e MultiExtractor) ExtractToken(ctx context.Context) (string, error) {
	// loop over extractors and return the first one that contains data
	for _, extractor := range e {
		if tok, err := extractor.ExtractToken(ctx); tok != "" {
			return tok, nil
		} else if !errors.Is(err, request.ErrNoTokenInRequest) {
			return "", err
		}
	}
	return "", request.ErrNoTokenInRequest
}

// PostExtractionFilter wraps an Extractor in this to post-process the value before it's handed off.
// See AuthorizationHeaderExtractor for an example
type PostExtractionFilter struct {
	Extractor
	Filter func(string) (string, error)
}

func (e *PostExtractionFilter) ExtractToken(ctx context.Context) (string, error) {
	if tok, err := e.Extractor.ExtractToken(ctx); tok != "" {
		return e.Filter(tok)
	} else {
		return "", err
	}
}

// MetadataExtractor extracts the JWT token from the gRPC request metadata using the passed in metadata key.
type MetadataExtractor string

func (e MetadataExtractor) String() string {
	return string(e)
}

func (e MetadataExtractor) ExtractToken(ctx context.Context) (string, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", request.ErrNoTokenInRequest
	}

	values := md[e.String()]
	if len(values) == 0 {
		return "", request.ErrNoTokenInRequest
	}

	value := values[0]
	if value == "" {
		return "", request.ErrNoTokenInRequest
	}

	return value, nil

}

type GenerateExtractor struct {
	generate func() (string, error)
}

func NewGenerateExtractor(
	generate func() (string, error),
) GenerateExtractor {
	return GenerateExtractor{
		generate: generate,
	}
}

func NewGenerateExtractorWithValue(token string) GenerateExtractor {
	return NewGenerateExtractor(
		func() (string, error) {
			return token, nil
		},
	)
}

func (e GenerateExtractor) ExtractToken(_ *http.Request) (string, error) {
	return e.generate()
}

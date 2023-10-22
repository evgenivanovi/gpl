package mw

import (
	"errors"
	"net"
	"net/http"
	"strings"

	"github.com/evgenivanovi/gpl/std"
	slices "github.com/evgenivanovi/gpl/std/slice"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
	"github.com/gookit/goutil/strutil"
)

// Errors
var (
	ErrNoIP      = errors.New("no ip present in request")
	ErrInvalidIP = errors.New("invalid ip")
)

type IPExtractor interface {
	ExtractIP(*http.Request) (net.IP, error)
}

type funcExtractor struct {
	extraction func(request *http.Request) (net.IP, error)
}

func (ext *funcExtractor) ExtractIP(
	request *http.Request,
) (net.IP, error) {
	return ext.extraction(request)
}

func NewIPExtractor(
	extraction func(*http.Request) (net.IP, error),
) IPExtractor {
	return &funcExtractor{
		extraction: extraction,
	}
}

// MultiIPExtractor tries IPExtractors in order until one returns a net.IP or an error occurs
type MultiIPExtractor []IPExtractor

func (ext MultiIPExtractor) ExtractIP(
	request *http.Request,
) (net.IP, error) {

	for _, extractor := range ext {

		ip, err := extractor.ExtractIP(request)

		if err == nil {
			return ip, nil
		}

		if !errors.Is(err, ErrNoIP) {
			return nil, err
		}

		if !errors.Is(err, ErrInvalidIP) {
			return nil, err
		}

	}

	return nil, ErrNoIP

}

type subnet struct {
	subnet     string
	extractors MultiIPExtractor

	recover func(http.ResponseWriter, *http.Request, error)
}

func NewSubnet(ops ...SubnetOp) func(next http.Handler) http.Handler {

	mw := defaultSubnet()

	for _, op := range ops {
		op(mw)
	}

	if strutil.IsNotBlank(mw.subnet) && slices.IsEmpty(mw.extractors) {
		panic("with subnet mask at least one ip extractor is required")
	}

	return mw.wrap

}

func defaultSubnet() *subnet {
	return &subnet{
		recover: func(w http.ResponseWriter, _ *http.Request, _ error) {
			w.WriteHeader(http.StatusUnauthorized)
		},
	}
}

func (s *subnet) wrap(next http.Handler) http.Handler {

	mwFunc := func(writer http.ResponseWriter, request *http.Request) {

		if strutil.IsBlank(s.subnet) {
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		ip, err := s.extractors.ExtractIP(request)
		if err != nil {
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		if isIPInSubnet(ip, s.subnet) {
			writer.WriteHeader(http.StatusForbidden)
			return
		}

		next.ServeHTTP(writer, request)

	}

	return http.HandlerFunc(mwFunc)

}

func isIPInSubnet(ip net.IP, subnet string) bool {
	ipv4 := ip.To4()

	_, ipnet, err := net.ParseCIDR(subnet)
	if err != nil {
		return false
	}

	return ipnet.Contains(ipv4)
}

type SubnetOp func(*subnet)

func WithSubnet(value string) SubnetOp {
	return func(s *subnet) {
		s.subnet = value
	}
}

func WithXRealIPExtractor() SubnetOp {

	extraction := func(request *http.Request) (net.IP, error) {

		ipexp := request.Header.Get(headers.XRealIP.String())
		if strutil.IsBlank(ipexp) {
			return nil, ErrNoIP
		}

		ip := net.ParseIP(ipexp)
		if ip == nil {
			return nil, ErrInvalidIP
		}

		return ip, nil

	}

	return func(s *subnet) {
		s.extractors = append(s.extractors, NewIPExtractor(extraction))
	}

}

func WithXForwardedForExtractor() SubnetOp {

	extraction := func(request *http.Request) (net.IP, error) {

		ips := request.Header.Get(headers.XForwardedFor.String())
		if strutil.IsBlank(ips) {
			return nil, ErrNoIP
		}

		ipexp := strings.Split(ips, std.Comma)[0]

		ip := net.ParseIP(ipexp)
		if ip == nil {
			return nil, ErrInvalidIP
		}

		return ip, nil

	}

	return func(s *subnet) {
		s.extractors = append(s.extractors, NewIPExtractor(extraction))
	}

}

package fw

import (
	"strconv"
	"strings"

	"github.com/evgenivanovi/gpl/std"
	"github.com/gookit/goutil/strutil"
)

/* __________________________________________________ */

const ServerDefaultHost = "localhost"
const ServerDefaultHTTPPort = 80
const ServerDefaultHTTPSPort = 443
const ServerDefaultGRPCPort = 82

/* __________________________________________________ */

type ServerOp func(*ServerSettings)

func (o ServerOp) Join(op ServerOp) ServerOp {
	return func(settings *ServerSettings) {
		o(settings)
		op(settings)
	}
}

func (o ServerOp) And(ops ...ServerOp) ServerOp {
	return func(settings *ServerSettings) {
		o(settings)
		for _, op := range ops {
			op(settings)
		}
	}
}

/* __________________________________________________ */

func WithHost(host string) ServerOp {
	return func(opts *ServerSettings) {
		opts.host = host
	}
}

func WithHostFn(fn func() string) ServerOp {
	return WithHost(fn())
}

/* __________________________________________________ */

func WithHttp(enabled bool) ServerOp {
	return func(settings *ServerSettings) {
		settings.httpEnabled = enabled
	}
}

func WithHttpEnabled() ServerOp {
	return func(settings *ServerSettings) {
		settings.httpEnabled = true
	}
}

func WithHttps(enabled bool) ServerOp {
	return func(settings *ServerSettings) {
		settings.httpsEnabled = enabled
	}
}

func WithHttpsEnabled() ServerOp {
	return func(settings *ServerSettings) {
		settings.httpsEnabled = true
	}
}

func WithGrpc(enabled bool) ServerOp {
	return func(settings *ServerSettings) {
		settings.grpcEnabled = enabled
	}
}

func WithGrpcEnabled() ServerOp {
	return func(settings *ServerSettings) {
		settings.grpcEnabled = true
	}
}

/* __________________________________________________ */

func WithStringHttpPort(port string) ServerOp {
	return func(settings *ServerSettings) {
		port, err := strutil.ToInt(port)
		if err != nil {
			panic(err)
		}
		settings.httpPort = port
	}
}

func WithHttpPort(port int) ServerOp {
	return func(settings *ServerSettings) {
		settings.httpPort = port
	}
}

func WithHttpPortFn(fn func() int) ServerOp {
	return WithHttpPort(fn())
}

/* __________________________________________________ */

func WithStringHttpsPort(port string) ServerOp {
	return func(settings *ServerSettings) {
		port, err := strutil.ToInt(port)
		if err != nil {
			panic(err)
		}
		settings.httpsPort = port
	}
}

func WithHttpsPort(port int) ServerOp {
	return func(opts *ServerSettings) {
		opts.httpsPort = port
	}
}

func WithHttpsPortFn(fn func() int) ServerOp {
	return WithHttpsPort(fn())
}

/* __________________________________________________ */

func WithStringGrpcPort(port string) ServerOp {
	return func(settings *ServerSettings) {
		port, err := strutil.ToInt(port)
		if err != nil {
			panic(err)
		}
		settings.grpcPort = port
	}
}

func WithGrpcPort(port int) ServerOp {
	return func(settings *ServerSettings) {
		settings.grpcPort = port
	}
}

func WithGrpcPortFn(fn func() int) ServerOp {
	return WithGrpcPort(fn())
}

/* __________________________________________________ */

func WithHttpAddress(address string) ServerOp {
	host, port, found := strings.Cut(address, std.Colon)

	if !found {
		return WithHost(address)
	}

	return WithHost(host).
		Join(WithStringHttpPort(port))
}

func WithHttpAddressFn(fn func() string) ServerOp {
	return WithHttpAddress(fn())
}

/* __________________________________________________ */

func WithHttpsAddress(address string) ServerOp {
	host, port, found := strings.Cut(address, std.Colon)

	if !found {
		return WithHost(address)
	}

	return WithHost(host).
		Join(WithStringHttpsPort(port))
}

func WithHttpsAddressFn(fn func() string) ServerOp {
	return WithHttpsAddress(fn())
}

/* __________________________________________________ */

func WithGrpcAddress(address string) ServerOp {
	host, port, found := strings.Cut(address, std.Colon)

	if !found {
		return WithHost(address)
	}

	return WithHost(host).
		Join(WithStringGrpcPort(port))
}

func WithGrpcAddressFn(fn func() string) ServerOp {
	return WithGrpcAddress(fn())
}

/* __________________________________________________ */

func WithTLS(fn func() *TLS) ServerOp {
	return func(opts *ServerSettings) {
		opts.tls = fn()
	}
}

/* __________________________________________________ */

type ServerSettings struct {
	host string

	httpPort    int
	httpEnabled bool

	httpsPort    int
	httpsEnabled bool

	grpcPort    int
	grpcEnabled bool

	tls *TLS
}

/* __________________________________________________ */

func (ss ServerSettings) HttpEnabled() bool {
	return ss.httpEnabled
}

func (ss ServerSettings) HttpsEnabled() bool {
	return ss.httpsEnabled
}

func (ss ServerSettings) GrpcEnabled() bool {
	return ss.grpcEnabled
}

/* __________________________________________________ */

func (ss ServerSettings) HttpAddress() string {
	res := strings.Builder{}
	res.WriteString(ss.host)
	res.WriteString(std.Colon)
	res.WriteString(strconv.Itoa(ss.httpPort))
	return res.String()
}

func (ss ServerSettings) HttpsAddress() string {
	res := strings.Builder{}
	res.WriteString(ss.host)
	res.WriteString(std.Colon)
	res.WriteString(strconv.Itoa(ss.httpsPort))
	return res.String()
}

func (ss ServerSettings) GrpcAddress() string {
	res := strings.Builder{}
	res.WriteString(ss.host)
	res.WriteString(std.Colon)
	res.WriteString(strconv.Itoa(ss.grpcPort))
	return res.String()
}

func NewServerSettings(ops ...ServerOp) *ServerSettings {
	settings := defaultServerSettings()
	for _, op := range ops {
		op(&settings)
	}
	return &settings
}

func defaultServerSettings() ServerSettings {
	return ServerSettings{
		host: ServerDefaultHost,

		httpPort:    ServerDefaultHTTPPort,
		httpEnabled: false,

		httpsPort:    ServerDefaultHTTPSPort,
		httpsEnabled: false,

		grpcPort:    ServerDefaultGRPCPort,
		grpcEnabled: false,

		tls: NewTLS(),
	}
}

/* __________________________________________________ */

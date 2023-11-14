package fw

import (
	"crypto/tls"
	"strconv"
	"strings"

	"github.com/evgenivanovi/gpl/std"
	"github.com/gookit/goutil/strutil"
)

/* __________________________________________________ */

const ServerDefaultHost = "localhost"
const ServerDefaultHTTPPort = 80
const ServerDefaultHTTPSPort = 443

/* __________________________________________________ */

type ServerOp func(*ServerSettings)

func (o ServerOp) Join(op ServerOp) ServerOp {
	return func(opts *ServerSettings) {
		o(opts)
		op(opts)
	}
}

func (o ServerOp) And(ops ...ServerOp) ServerOp {
	return func(opts *ServerSettings) {
		o(opts)
		for _, fn := range ops {
			fn(opts)
		}
	}
}

/* __________________________________________________ */

type ServerSettings struct {
	host string

	httpPort    int
	httpEnabled bool

	httpsPort    int
	httpsEnabled bool

	tls *TLS
}

/* __________________________________________________ */

func (op ServerSettings) HttpEnabled() bool {
	return op.httpEnabled
}

func (op ServerSettings) HttpsEnabled() bool {
	return op.httpsEnabled
}

/* __________________________________________________ */

func (op ServerSettings) HttpAddress() string {
	result := strings.Builder{}
	result.WriteString(op.host)
	result.WriteString(std.Colon)
	result.WriteString(strconv.Itoa(op.httpPort))
	return result.String()
}

func (op ServerSettings) HttpsAddress() string {
	result := strings.Builder{}
	result.WriteString(op.host)
	result.WriteString(std.Colon)
	result.WriteString(strconv.Itoa(op.httpsPort))
	return result.String()
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
	return func(opts *ServerSettings) {
		opts.httpEnabled = enabled
	}
}

func WithHttpEnabled() ServerOp {
	return func(opts *ServerSettings) {
		opts.httpEnabled = true
	}
}

func WithHttps(enabled bool) ServerOp {
	return func(opts *ServerSettings) {
		opts.httpsEnabled = enabled
	}
}

func WithHttpsEnabled() ServerOp {
	return func(opts *ServerSettings) {
		opts.httpsEnabled = true
	}
}

/* __________________________________________________ */

func WithStringHttpPort(port string) ServerOp {
	return func(opts *ServerSettings) {
		p, err := strutil.ToInt(port)
		if err != nil {
			panic(err)
		}
		opts.httpPort = p
	}
}

func WithHttpPort(port int) ServerOp {
	return func(opts *ServerSettings) {
		opts.httpPort = port
	}
}

func WithHttpPortFn(fn func() int) ServerOp {
	return WithHttpPort(fn())
}

/* __________________________________________________ */

func WithStringHttpsPort(port string) ServerOp {
	return func(opts *ServerSettings) {
		p, err := strutil.ToInt(port)
		if err != nil {
			panic(err)
		}
		opts.httpsPort = p
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

func WithHttpAddress(address string) ServerOp {
	host, port, found := strings.Cut(address, std.Colon)
	if !found {
		return WithHost(address)
	}
	return WithHost(host).Join(WithStringHttpPort(port))
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
	return WithHost(host).Join(WithStringHttpsPort(port))
}

func WithHttpsAddressFn(fn func() string) ServerOp {
	return WithHttpsAddress(fn())
}

/* __________________________________________________ */

func WithTLS(fn func() *TLS) ServerOp {
	return func(opts *ServerSettings) {
		opts.tls = fn()
	}
}

/* __________________________________________________ */

func NewServerSettings(opts ...ServerOp) *ServerSettings {
	settings := defaultServerSettings()
	for _, op := range opts {
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

		tls: NewTLS(),
	}
}

/* __________________________________________________ */

type TLSOp func(*TLS)

func WithConfig(config *tls.Config) TLSOp {
	return func(tls *TLS) {
		if tls != nil {
			tls.config = config
		}
	}
}

func WithCert(cert string) TLSOp {
	return func(tls *TLS) {
		if tls != nil {
			tls.cert = cert
		}
	}
}

func WithKey(key string) TLSOp {
	return func(tls *TLS) {
		if tls != nil {
			tls.key = key
		}
	}
}

func WithCertKey(cert, key string) TLSOp {
	return func(tls *TLS) {
		tls.cert = cert
		tls.key = key
	}
}

/* __________________________________________________ */

type TLS struct {
	config *tls.Config
	cert   string
	key    string
}

func (tls *TLS) Enabled() bool {
	return tls.EnabledAutoTLS() || tls.EnabledNonAutoTLS()
}

func (tls *TLS) Disabled() bool {
	return !tls.Enabled()
}

// EnabledAutoTLS
// Filenames containing a certificate and matching private key for the
// server must be provided if neither the Server's TLSConfig.Certificates
// nor TLSConfig.GetCertificate are populated.
func (tls *TLS) EnabledAutoTLS() bool {

	if tls == nil || tls.config == nil {
		return false
	}

	if len(tls.config.Certificates) == 0 || tls.config.GetCertificate == nil {
		return false
	}

	return strutil.IsBlank(tls.cert) && strutil.IsBlank(tls.key)

}

// EnabledNonAutoTLS
// Filenames containing a certificate and matching private key for the
// server must be provided if neither the Server's TLSConfig.Certificates
// nor TLSConfig.GetCertificate are populated.
func (tls *TLS) EnabledNonAutoTLS() bool {

	if tls == nil || tls.config == nil {
		return false
	}

	if len(tls.config.Certificates) != 0 || tls.config.GetCertificate != nil {
		return false
	}

	return strutil.IsNotBlank(tls.cert) && strutil.IsNotBlank(tls.key)

}

func NewTLS(opts ...TLSOp) *TLS {
	cfg := &TLS{}
	for _, fn := range opts {
		fn(cfg)
	}
	return cfg
}

/* __________________________________________________ */

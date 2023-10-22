package fw

import (
	"crypto/tls"

	"github.com/gookit/goutil/strutil"
)

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

func NewTLS(ops ...TLSOp) *TLS {
	cfg := &TLS{}
	for _, op := range ops {
		op(cfg)
	}
	return cfg
}

package fw

import (
	"net/http"

	grpcfw "github.com/evgenivanovi/gpl/server"
	"google.golang.org/grpc"
)

/* __________________________________________________ */

type ConfigurationOp func(*Configuration)

func WithApplication(app *Application) ConfigurationOp {
	return func(cfg *Configuration) {
		cfg.App = app
	}
}

func WithHTTPHandler(handler http.Handler) ConfigurationOp {
	return func(cfg *Configuration) {
		cfg.HTTPHandler = handler
	}
}

func WithGRPCServices(services ...grpcfw.GRPCService) ConfigurationOp {
	return func(cfg *Configuration) {
		cfg.GRPCServices = append(cfg.GRPCServices, services...)
	}
}

func WithGRPCUnaryMW(mws ...grpc.UnaryServerInterceptor) ConfigurationOp {
	return func(cfg *Configuration) {
		cfg.GRPCUnaryMWs = append(cfg.GRPCUnaryMWs, mws...)
	}
}

func WithGRPCStreamMW(mws ...grpc.StreamServerInterceptor) ConfigurationOp {
	return func(cfg *Configuration) {
		cfg.GRPCStreamMWs = append(cfg.GRPCStreamMWs, mws...)
	}
}

/* __________________________________________________ */

type Configuration struct {
	App *Application

	HTTPHandler http.Handler

	GRPCServices  []grpcfw.GRPCService
	GRPCUnaryMWs  []grpc.UnaryServerInterceptor
	GRPCStreamMWs []grpc.StreamServerInterceptor
}

func NewConfiguration(ops ...ConfigurationOp) *Configuration {
	cfg := defaultConfiguration()
	for _, op := range ops {
		op(cfg)
	}
	return cfg
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		App:           NewApplication(),
		GRPCServices:  make([]grpcfw.GRPCService, 0),
		GRPCUnaryMWs:  make([]grpc.UnaryServerInterceptor, 0),
		GRPCStreamMWs: make([]grpc.StreamServerInterceptor, 0),
	}
}

func (c *Configuration) WithHTTPHandler(handler http.Handler) {
	if c != nil {
		c.HTTPHandler = handler
	}
}

func (c *Configuration) WithGRPCServices(services ...grpcfw.GRPCService) {
	if c != nil && len(services) != 0 {
		c.GRPCServices = append(c.GRPCServices, services...)
	}
}

func (c *Configuration) WithGRPCUnaryMW(mws ...grpc.UnaryServerInterceptor) {
	if c != nil && len(mws) != 0 {
		c.GRPCUnaryMWs = append(c.GRPCUnaryMWs, mws...)
	}
}

func (c *Configuration) WithGRPCStreamMW(mws ...grpc.StreamServerInterceptor) {
	if c != nil && len(mws) != 0 {
		c.GRPCStreamMWs = append(c.GRPCStreamMWs, mws...)
	}
}

/* __________________________________________________ */

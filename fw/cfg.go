package fw

import (
	"net/http"

	grpcfw "github.com/evgenivanovi/gpl/server"
	"google.golang.org/grpc"
)

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

type Configuration struct {
	App *Application

	HTTPHandler http.Handler

	GRPCReflection bool
	GRPCServices   []grpcfw.GRPCService
	GRPCUnaryMWs   []grpc.UnaryServerInterceptor
	GRPCStreamMWs  []grpc.StreamServerInterceptor
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

func (c *Configuration) WithHTTPHandler(handler http.Handler) *Configuration {
	if c != nil {
		c.HTTPHandler = handler
	}
	return c
}

func (c *Configuration) WithGRPCReflection(reflection bool) *Configuration {
	if c != nil {
		c.GRPCReflection = reflection
	}
	return c
}

func (c *Configuration) WithGRPCServices(services ...grpcfw.GRPCService) *Configuration {
	if c != nil && len(services) != 0 {
		c.GRPCServices = append(c.GRPCServices, services...)
	}
	return c
}

func (c *Configuration) WithGRPCUnaryMW(mws ...grpc.UnaryServerInterceptor) *Configuration {
	if c != nil && len(mws) != 0 {
		c.GRPCUnaryMWs = append(c.GRPCUnaryMWs, mws...)
	}
	return c
}

func (c *Configuration) WithGRPCStreamMW(mws ...grpc.StreamServerInterceptor) *Configuration {
	if c != nil && len(mws) != 0 {
		c.GRPCStreamMWs = append(c.GRPCStreamMWs, mws...)
	}
	return c
}

package server

import (
	"context"
	"log/slog"
	"net"
	"strconv"

	netx "github.com/evgenivanovi/gpl/stdx/net"
	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCService interface {
	RegisterService(grpc.ServiceRegistrar)
}

// GRPCServerOp - is callback function that applies an option to GrpcServer.
type GRPCServerOp func(*GRPCServer)

// WithGrpcServerConfig - adds GrpcServerConfig to GrpcServer.
func WithGrpcServerConfig(config GRPCServerConfig) GRPCServerOp {
	return func(server *GRPCServer) {
		server.config = config
	}
}

// WithGrpcReflection - adds or removes reflection to GrpcServer.
func WithGrpcReflection(reflection bool) GRPCServerOp {
	return func(server *GRPCServer) {
		server.reflection = reflection
	}
}

// WithService - adds GRPCService to GrpcServer.
func WithService(service GRPCService) GRPCServerOp {
	return func(server *GRPCServer) {
		services := make([]GRPCService, 0)
		services = append(services, service)
		server.services = services
	}
}

// WithServices - add []GRPCService to GrpcServer.
func WithServices(services ...GRPCService) GRPCServerOp {
	return func(server *GRPCServer) {
		server.services = services
	}
}

// AddService - appends GRPCService to GrpcServer.
func AddService(service GRPCService) GRPCServerOp {
	return func(server *GRPCServer) {
		server.services = append(server.services, service)
	}
}

// AddServices - append []GRPCService to GrpcServer.
func AddServices(services ...GRPCService) GRPCServerOp {
	return func(server *GRPCServer) {
		server.services = append(server.services, services...)
	}
}

// WithStreamInterceptors - add []grpc.StreamServerInterceptor to GrpcServer.
func WithStreamInterceptors(in ...grpc.StreamServerInterceptor) GRPCServerOp {
	return func(server *GRPCServer) {
		server.streamInterceptors = append(server.streamInterceptors, in...)
	}
}

// WithUnaryInterceptors - add []grpc.UnaryServerInterceptor to GrpcServer.
func WithUnaryInterceptors(in ...grpc.UnaryServerInterceptor) GRPCServerOp {
	return func(server *GRPCServer) {
		server.unaryInterceptors = append(server.unaryInterceptors, in...)
	}
}

// WithLogger - adds slog.Logger to GrpcServer.
func WithLogger(log *slog.Logger) GRPCServerOp {
	return func(server *GRPCServer) {
		server.log = log
	}
}

type GRPCServerConfig struct {
	Port int
}

func NewGRPCServerConfig(port int) *GRPCServerConfig {
	return &GRPCServerConfig{
		Port: port,
	}
}

func (cfg GRPCServerConfig) PortString() string {
	return strconv.Itoa(cfg.Port)
}

type GRPCServer struct {
	server             *grpc.Server
	reflection         bool
	services           []GRPCService
	config             GRPCServerConfig
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor

	log *slog.Logger
}

// NewGrpcServer - creates new GrpcServer with options via provided GrpcServerOption.
func NewGrpcServer(ops ...GRPCServerOp) *GRPCServer {
	server := &GRPCServer{}

	for _, op := range ops {
		op(server)
	}

	return server
}

// RegisterServices - adds a service to gRPC server via,
// GRPCService.RegisterService function which is called on each provided GRPCService.
func (s *GRPCServer) RegisterServices(services ...GRPCService) {
	for _, service := range services {
		service.RegisterService(s.server)
	}
}

// StartChannable - starts gRPC server.
func (s *GRPCServer) StartChannable(
	cancel chan error,
) {
	onError := func(err error) {
		cancel <- err
	}
	s.executeStart(onError)
}

// StartCancellable - starts gRPC server.
func (s *GRPCServer) StartCancellable(
	cancel context.CancelFunc,
) {
	onError := func(error) {
		cancel()
	}
	s.executeStart(onError)
}

func (s *GRPCServer) executeStart(onError func(error)) {

	conn, err := net.Listen(netx.TCP, ":"+s.config.PortString())
	if err != nil {
		onError(err)
		return
	}

	s.server = grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmw.ChainStreamServer(
				s.streamInterceptors...,
			),
		),
		grpc.UnaryInterceptor(
			grpcmw.ChainUnaryServer(
				s.unaryInterceptors...,
			),
		),
	)

	s.RegisterServices(s.services...)

	if s.reflection {
		reflection.Register(s.server)
	}

	go func() {
		err = s.server.Serve(conn)
		if err != nil {
			onError(err)
			return
		}
	}()

}

// Stop - gracefully stops gRPC server.
func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
}

package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type UnaryServerErrorInterceptor func(
	ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, ex error,
) (resp any, err error)

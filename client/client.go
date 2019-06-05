package client

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// DefaultDialOptions returns default gRPC dial options.
func DefaultDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
			grpc_opentracing.StreamClientInterceptor(),
			grpc_prometheus.StreamClientInterceptor,
			grpc_retry.StreamClientInterceptor(),
		)),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			grpc_opentracing.UnaryClientInterceptor(),
			grpc_prometheus.UnaryClientInterceptor,
			grpc_retry.UnaryClientInterceptor(),
		)),
	}
}

// New returns a new gRPC client connection.
func New(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if opts == nil {
		opts = DefaultDialOptions()
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		err = errors.Wrapf(err, "failed to create gRPC client connection to the given target: %v", target)
	}
	return conn, err
}

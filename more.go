package more

import (
	"context"

	"google.golang.org/grpc"
)

// RunServer runs an entrypoint of the grpc server.
func RunServer(ctx context.Context, register func(*grpc.Server)) error {
	return runServer(ctx, register)
}

// RunGateway runs an entrypoint of the proxy server.
func RunGateway(ctx context.Context, registers ...registerHandlerFromEndpointFunc) error {
	return runGateway(ctx, registers...)
}

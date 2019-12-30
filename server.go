package more

import (
	"context"
	"net"

	"github.com/gotopia/more/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func runServer(ctx context.Context, register func(*grpc.Server)) error {
	network := config.Server.Network()
	address := config.Server.Address()
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			err := errors.Wrapf(err, "failed to close %v %v", network, address)
			zap.L().Error(err.Error())
		}
	}()

	s := newServer()
	register(s)
	grpc_prometheus.Register(s)

	go func() {
		defer func() {
			s.GracefulStop()
		}()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

func newServer() *grpc.Server {
	registerDefaultInterceptors()
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamInterceptors()...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryInterceptors()...,
		)),
	)
}

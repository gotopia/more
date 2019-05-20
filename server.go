package more

import (
	"context"
	"net"

	"github.com/gotopia/more/auth"
	"github.com/gotopia/more/config"
	"github.com/gotopia/more/recovery"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
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
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(zap.L()),
			grpc_zap.PayloadStreamServerInterceptor(zap.L(), payloadLoggingDecider),
			grpc_auth.StreamServerInterceptor(auth.Func),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recovery.Handler)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(zap.L()),
			grpc_zap.PayloadUnaryServerInterceptor(zap.L(), payloadLoggingDecider),
			grpc_auth.UnaryServerInterceptor(auth.Func),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recovery.Handler)),
		)),
	)
}

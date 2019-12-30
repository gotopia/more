package more

import (
	"github.com/gotopia/more/auth"
	"github.com/gotopia/more/config"
	"github.com/gotopia/more/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

type interceptor struct {
	identifier              string
	streamServerInterceptor grpc.StreamServerInterceptor
	unaryServerInterceptor  grpc.UnaryServerInterceptor
}

var interceptorMap map[string]*interceptor

func init() {
	interceptorMap = make(map[string]*interceptor)
}

func registerDefaultInterceptors() {
	registerInterceptor("ctxtags", grpc_ctxtags.StreamServerInterceptor(), grpc_ctxtags.UnaryServerInterceptor())
	registerInterceptor("opentracing", grpc_opentracing.StreamServerInterceptor(), grpc_opentracing.UnaryServerInterceptor())
	registerInterceptor("prometheus", grpc_prometheus.StreamServerInterceptor, grpc_prometheus.UnaryServerInterceptor)
	registerInterceptor("zap", grpc_zap.StreamServerInterceptor(zap.L()), grpc_zap.UnaryServerInterceptor(zap.L()))
	registerInterceptor("payload", grpc_zap.PayloadStreamServerInterceptor(zap.L(), payloadLoggingDecider), grpc_zap.PayloadUnaryServerInterceptor(zap.L(), payloadLoggingDecider))
	registerInterceptor("auth", grpc_auth.StreamServerInterceptor(auth.Func), grpc_auth.UnaryServerInterceptor(auth.Func))
	registerInterceptor("validator", grpc_validator.StreamServerInterceptor(), grpc_validator.UnaryServerInterceptor())
	registerInterceptor("recovery", grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recovery.Handler)), grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recovery.Handler)))
}

func registerInterceptor(identifier string, streamServerInterceptor grpc.StreamServerInterceptor, unaryServerInterceptor grpc.UnaryServerInterceptor) {
	interceptorMap[identifier] = &interceptor{
		streamServerInterceptor: streamServerInterceptor,
		unaryServerInterceptor:  unaryServerInterceptor,
	}
}

func streamInterceptors() []grpc.StreamServerInterceptor {
	identifiers := config.Server.Interceptors()
	streamInterceptors := []grpc.StreamServerInterceptor{}
	for _, identifier := range identifiers {
		interceptor, ok := interceptorMap[identifier]
		if ok {
			streamInterceptors = append(streamInterceptors, interceptor.streamServerInterceptor)
		}
	}
	return streamInterceptors
}

func unaryInterceptors() []grpc.UnaryServerInterceptor {
	identifiers := config.Server.Interceptors()
	unaryInterceptors := []grpc.UnaryServerInterceptor{}
	for _, identifier := range identifiers {
		interceptor, ok := interceptorMap[identifier]
		if ok {
			unaryInterceptors = append(unaryInterceptors, interceptor.unaryServerInterceptor)
		}
	}
	return unaryInterceptors
}

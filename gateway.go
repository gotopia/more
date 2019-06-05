package more

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gotopia/more/client"
	"github.com/gotopia/more/config"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

type registerHandlerFromEndpointFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func runGateway(ctx context.Context, registers ...registerHandlerFromEndpointFunc) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	gmux := runtime.NewServeMux(
		runtime.WithMetadata(injectHeadersIntoMetadata),
	)
	for _, register := range registers {
		if err := register(ctx, gmux, config.Server.Address(), client.DefaultDialOptions()); err != nil {
			return errors.Wrap(err, "failed to register handler from end point")
		}
	}
	mux.Handle("/", gmux)

	httpHandler := handlers.CORS(
		handlers.AllowedOrigins(config.Cors.Origins()),
		handlers.AllowedMethods(config.Cors.Methods()),
		handlers.AllowedHeaders(config.Cors.Headers()),
	)(mux)

	return http.ListenAndServe(config.Gateway.Address(), httpHandler)
}

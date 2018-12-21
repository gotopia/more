package more

import (
	"context"

	"github.com/gotopia/more/config"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
)

func init() {
	var logger *zap.Logger
	if config.Development() {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	grpc_zap.ReplaceGrpcLogger(logger)
	zap.ReplaceGlobals(logger)
}

func payloadLoggingDecider(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
	return !config.Development()
}

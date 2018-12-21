package recovery

import (
	"github.com/gotopia/more/config"
	"github.com/gotopia/sin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func internalToStatus(err error) *status.Status {
	s := status.New(codes.Internal, "Internal server error.")
	if config.Development() {
		i := sin.NewDebugInfo(err)
		s, _ = s.WithDetails(i.Serialize())
	}
	return s
}

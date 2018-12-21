package recovery

import (
	"github.com/gotopia/insight/parser"
	"github.com/gotopia/sin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parsingErrorToStatus(err parser.ParsingError) *status.Status {
	s := status.New(codes.InvalidArgument, "Invalid Argument.")
	fv := sin.NewFieldViolation("filter", err.Error())
	r := sin.NewBadRequest(fv)
	s, _ = s.WithDetails(r.Serialize())
	return s
}

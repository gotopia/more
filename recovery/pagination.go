package recovery

import (
	"github.com/gotopia/ripper"
	"github.com/gotopia/sin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func paginationErrorToStatus(err ripper.PaginationError) *status.Status {
	s := status.New(codes.InvalidArgument, "Invalid Argument.")
	field := err.Field()
	fv := sin.NewFieldViolation(field, err.Error())
	r := sin.NewBadRequest(fv)
	s, _ = s.WithDetails(r.Serialize())
	return s
}

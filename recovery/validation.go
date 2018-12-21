package recovery

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/gotopia/sin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validationErrorsToStatus(errs validation.Errors) *status.Status {
	s := status.New(codes.InvalidArgument, "Invalid Argument.")
	for field, err := range errs {
		fv := sin.NewFieldViolation(field, err.Error())
		r := sin.NewBadRequest(fv)
		s, _ = s.WithDetails(r.Serialize())
	}
	return s
}

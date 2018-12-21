package recovery

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-sql-driver/mysql"
	"github.com/gotopia/insight/parser"
	"github.com/gotopia/ripper"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is default recovery handler.
func Handler(p interface{}) error {
	if err, ok := p.(error); ok {
		return handleError(err)
	}
	return status.New(codes.Unknown, "Internal server error.").Err()
}

func handleError(err error) error {
	cause := errors.Cause(err)
	s, ok := status.FromError(cause)
	if ok {
		return s.Err()
	}
	zap.L().Info("error logged as error field", zap.Error(cause))
	switch cause {
	case sql.ErrNoRows:
		return status.New(codes.NotFound, "Not found.").Err()
	default:
		return handleTypedError(err)
	}
}

func handleTypedError(err error) error {
	switch te := errors.Cause(err).(type) {
	case *jwt.ValidationError:
		return status.New(codes.Unauthenticated, te.Error()).Err()
	case *mysql.MySQLError:
		return mySQLErrorToStatus(te).Err()
	case validation.Errors:
		return validationErrorsToStatus(te).Err()
	case ripper.PaginationError:
		return paginationErrorToStatus(te).Err()
	case parser.ParsingError:
		return parsingErrorToStatus(te).Err()
	default:
		return internalToStatus(err).Err()
	}
}

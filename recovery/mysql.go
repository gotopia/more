package recovery

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gotopia/sin"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var msgMySQL1062 = regexp.MustCompile("Duplicate entry '([^']+)' for key '([^']+)'")

func mySQLErrorToStatus(err *mysql.MySQLError) *status.Status {
	msg := err.Message
	switch err.Number {
	case 1062:
		subs := msgMySQL1062.FindStringSubmatch(msg)
		if len(subs) != 3 {
			return internalToStatus(err)
		}
		entry := subs[1]
		key := subs[2]
		idx := strings.TrimPrefix(key, "idx_")
		s := status.New(codes.AlreadyExists, "Already exists.")
		fv := sin.NewFieldViolation(idx, fmt.Sprintf("'%v' has already been taken", entry))
		r := sin.NewBadRequest(fv)
		s, _ = s.WithDetails(r.Serialize())
		return s
	default:
		return internalToStatus(err)
	}
}

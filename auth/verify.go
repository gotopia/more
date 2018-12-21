package auth

import (
	"context"
	"strconv"

	"github.com/gotopia/more/config"
	"github.com/gotopia/watcher"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

var authKey = struct{}{}

// Func uses the user-specified auth function.
func Func(ctx context.Context) (context.Context, error) {
	switch config.Auth.Type() {
	case "jwt":
		return jwtAuthFunc(ctx)
	default:
		return ctx, nil
	}
}

func jwtAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		panic(err)
	}
	sub, err := watcher.Verify(token, config.Auth.Issuer())
	if err != nil {
		panic(err)
	}
	grpc_ctxtags.Extract(ctx).Set("auth.sub", sub)
	ctx = context.WithValue(ctx, authKey, sub)
	return ctx, err
}

// CurrentUserID retrieves the current user_id from context. Only works when sub is an integer.
func CurrentUserID(ctx context.Context) uint {
	sub, _ := CurrentSub(ctx)
	userID, _ := strconv.ParseUint(sub, 10, 32)
	return uint(userID)
}

// CurrentSub retrieves the current sub from context.
func CurrentSub(ctx context.Context) (sub string, ok bool) {
	sub, ok = ctx.Value(authKey).(string)
	return
}

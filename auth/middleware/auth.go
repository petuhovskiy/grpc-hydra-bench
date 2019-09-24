package middleware

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/petuhovskiy/grpc-hydra-bench/auth/xctx"
)

type Auth struct {
	Validate func(token string) (interface{}, error)
}

func (a *Auth) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var authorization []string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		authorization = md["authorization"]
	}

	if len(authorization) > 1 {
		return nil, status.Errorf(codes.InvalidArgument, "too many authorizantion headers")
	}

	if len(authorization) == 1 {
		token := authorization[0]
		token = strings.TrimPrefix(token, "Bearer ")

		res, err := a.Validate(token)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, xctx.Token, res)
	}

	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

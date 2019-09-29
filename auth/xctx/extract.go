package xctx

import (
	"context"

	"github.com/ory/hydra/sdk/go/hydra/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrUnauthenticated = status.Errorf(codes.PermissionDenied, "user not authorized")

func GetInrospectedToken(ctx context.Context) (*models.Introspection, error) {
	token, ok := ctx.Value(Token).(*models.Introspection)
	if !ok {
		return nil, ErrUnauthenticated
	}

	return token, nil
}

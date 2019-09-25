package xctx

import (
	"context"

	"github.com/petuhovskiy/grpc-hydra-bench/auth/hydra/client/admin"
)

func GetInrospectedToken(ctx context.Context) *admin.IntrospectOAuth2TokenOK {
	return ctx.Value(Token).(*admin.IntrospectOAuth2TokenOK)
}

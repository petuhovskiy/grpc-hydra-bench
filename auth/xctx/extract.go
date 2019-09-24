package xctx

import (
	"context"

	"github.com/petuhovskiy/grpc-hydra-bench/hydracli/client/admin"
)

func GetInrospectedToken(ctx context.Context) *admin.IntrospectOAuth2TokenOK {
	return ctx.Value(Token).(*admin.IntrospectOAuth2TokenOK)
}

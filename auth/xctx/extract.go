package xctx

import (
	"context"

	"github.com/ory/hydra/sdk/go/hydra/client/admin"
)

func GetInrospectedToken(ctx context.Context) *admin.IntrospectOAuth2TokenOK {
	return ctx.Value(Token).(*admin.IntrospectOAuth2TokenOK)
}

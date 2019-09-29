package libauth

import (
	"strings"

	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInactiveToken = status.Errorf(codes.PermissionDenied, "token is inactive")

func Validator(cli *admin.Client, scopes ...string) func(string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		var scope *string
		if len(scopes) > 0 {
			tmp := strings.Join(scopes, " ")
			scope = &tmp
		}
		params := &admin.IntrospectOAuth2TokenParams{
			Scope: scope,
			Token: token,
		}

		resp, err := cli.IntrospectOAuth2Token(params, nil)
		if err != nil {
			return nil, err
		}

		payload := resp.Payload
		if payload.Active == nil || *payload.Active == false {
			return nil, ErrInactiveToken
		}

		return payload, nil
	}
}

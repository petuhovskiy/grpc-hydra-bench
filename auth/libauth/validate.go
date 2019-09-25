package libauth

import (
	"strings"

	"github.com/petuhovskiy/grpc-hydra-bench/auth/hydra/client/admin"
)

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

		return cli.IntrospectOAuth2Token(params, nil)
	}
}

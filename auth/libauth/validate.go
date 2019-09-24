package libauth

import (
	"strings"

	"github.com/petuhovskiy/grpc-hydra-bench/hydracli/client/admin"
)

func Validator(cli *admin.Client, scopes ...string) func(string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		params := admin.IntrospectOAuth2TokenParams{
			Scope: strings.Join(scopes, " "),
			Token: token,
		}

		return cli.IntrospectOAuth2Token(params, nil)
	}
}

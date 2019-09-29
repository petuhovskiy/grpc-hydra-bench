package hydracon

import (
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/ory/hydra/sdk/go/hydra/models"
)

func acceptLogin(adm *admin.Client, challenge string, subject string) (string, error) {
	req, err := adm.AcceptLoginRequest(
		admin.NewAcceptLoginRequestParams().
			WithLoginChallenge(challenge).
			WithBody(&models.HandledLoginRequest{
				Subject: &subject,
			}),
	)

	var redirectTo string
	if req != nil {
		if r2 := req.Payload; r2 != nil {
			redirectTo = r2.RedirectTo
		}
	}

	return redirectTo, err
}

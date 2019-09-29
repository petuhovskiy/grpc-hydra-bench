package hydracon

import (
	"fmt"
	"net/http"

	"github.com/ory/hydra/sdk/go/hydra/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/petuhovskiy/grpc-hydra-bench/auth/users"
)

type Handler struct {
	repo *users.Repo
	adm  *admin.Client
}

func NewHandler(repo *users.Repo, adm *admin.Client) *Handler {
	return &Handler{
		repo: repo,
		adm:  adm,
	}
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	spew.Dump(r.Cookies())

	if r.Method == "GET" {
		h.loginGET(w, r)
		return
	}

	h.loginPOST(w, r)
}

func (h *Handler) loginGET(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("login_challenge")

	login, err := h.adm.GetLoginRequest(
		admin.NewGetLoginRequestParams().
			WithLoginChallenge(challenge),
	)
	if err != nil {
		onError(err, w)
		return
	}

	payload := login.Payload
	if payload.Skip {
		red, err := acceptLogin(h.adm, challenge, payload.Subject)
		if err != nil {
			onError(err, w)
			return
		}

		http.Redirect(w, r, red, http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, loginTemplate, challenge)
}

func (h *Handler) loginPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	form := r.PostForm

	challenge := form.Get("challenge")
	username := form.Get("username")

	user, err := h.repo.FindByUsername(username)
	if err != nil {
		onError(err, w)
		return
	}

	subject := fmt.Sprintf("%v", user.ID)
	red, err := acceptLogin(h.adm, challenge, subject)
	if err != nil {
		onError(err, w)
		return
	}

	http.Redirect(w, r, red, http.StatusTemporaryRedirect)
	return
}

func (h *Handler) consent(w http.ResponseWriter, r *http.Request) {
	challenge := r.URL.Query().Get("consent_challenge")

	req, err := h.adm.GetConsentRequest(
		admin.NewGetConsentRequestParams().
			WithConsentChallenge(challenge),
	)
	if err != nil {
		onError(err, w)
		return
	}

	payload := req.Payload

	// subject := payload.Subject

	ac, err := h.adm.AcceptConsentRequest(
		admin.NewAcceptConsentRequestParams().
			WithConsentChallenge(challenge).
			WithBody(&models.HandledConsentRequest{
				GrantedAudience: payload.RequestedAudience,
				GrantedScope:    payload.RequestedScope,
				Session: &models.ConsentRequestSessionData{
					AccessToken: map[string]interface{}{
						//"username": "123", // TODO:
					},
				},
			}),
	)
	if err != nil {
		onError(err, w)
		return
	}

	http.Redirect(w, r, ac.Payload.RedirectTo, http.StatusTemporaryRedirect)
	return
}

func (h *Handler) createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", h.login)
	mux.HandleFunc("/consent", h.consent)
	return mux
}

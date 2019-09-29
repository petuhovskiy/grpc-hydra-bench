package hydracon

import (
	"net/http"

	"github.com/urfave/negroni"
)

func NewLoginAndConsentServer(handler *Handler) *http.Server {
	mux := handler.createMux()

	// CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))

	n := negroni.Classic()
	n.UseHandler(
		// CSRF(mux),
		mux,
	)

	return &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
}

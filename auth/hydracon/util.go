package hydracon

import (
	"log"
	"net/http"
)

func onError(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

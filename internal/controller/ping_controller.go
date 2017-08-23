package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	LogAndWriteResponse(w, r, func(l *log.Entry) ([]byte, int) {

		err := serviceAvailabilityCheck()
		if err != nil {
			return []byte(err.Error()), http.StatusInternalServerError
		}

		return []byte(`{"Ping":"Pong"}`), http.StatusOK
	})
}

func serviceAvailabilityCheck() error {
	//Do operation here that ensure that your service is completely healthy
	//TODO
	return nil
}

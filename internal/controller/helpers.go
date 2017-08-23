package controller

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/maknahar/go-web-skeleton/internal/utils/logger"
)

func writeHTTPResponse(response []byte, statusCode int, w http.ResponseWriter, start time.Time, l *log.Logger) {
	l.Printf("Status Code: %d, Response time: %f, Response: %s", statusCode, time.Since(start).Seconds(), string(response))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func LogAndWriteResponse(w http.ResponseWriter, r *http.Request, h func(l *log.Entry) ([]byte, int)) {
	start := time.Now()
	l := logger.GetLogger()

	response, statusCode := h(l)

	if statusCode >= 400 {
		response = jsonifyErrMessage(string(response), statusCode, nil)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(response)
	l.Printf("Status Code: %d, Response time: %v, Response: %s", statusCode, time.Since(start), string(response))

}

func jsonifyErrMessage(errMsg string, code int, err error) []byte {
	var errObj struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	errObj.Message = errMsg
	if err != nil {
		errObj.Message += " " + err.Error()
	}

	errObj.Code = code
	errObj.Status = "ERROR"
	data, _ := json.Marshal(errObj)
	return data
}

type SlydesErrObj struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/maknahar/go-web-skeleton/internal/utils/config"
	"github.com/maknahar/go-web-skeleton/internal/utils/logger"
)

//StartDMSPinger is a goroutine to keep pinging Dead Man Snitch.
func StartDMSPinger() {
	dmsURL := config.DMSURL
	l := logger.GetLogger()
	l.Println("DMS Url:", dmsURL)
	if dmsURL != "" {
		l.Println("Starting DMS Pinger")
		for {
			time.Sleep(time.Minute)
			res, err := http.Get(dmsURL + "?m=jt-greetingd+service+is+up") // To Check
			if err != nil {
				log.Println("DMS Ping Failed. ", err)
				continue
			}
			d, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Println("DMS Response Read Failed. ", err)
				continue
			}
			log.Println("DMS Response: ", string(d))
			time.Sleep(time.Minute)
		}
	}
	l.Println("DMS is not configured!!!")
}

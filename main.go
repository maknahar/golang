package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RealImage/que-ingester/internal/router"
	"github.com/RealImage/que-ingester/internal/utils/config"
	"github.com/RealImage/que-ingester/internal/utils/logger"
)

func main() {
	l := logger.GetLogger()

	l.Infoln("Configuring the app...")
	err := config.Initialize()
	if err != nil {
		l.Fatal(err)
	}

	l.Infoln("Starting Web Server on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router.GetRoutes()))
}

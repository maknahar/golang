package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/maknahar/go-web-skeleton/internal/router"
	"github.com/maknahar/go-web-skeleton/internal/utils/config"
)

func main() {
	fmt.Println("Configuring the app...")
	err := config.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting Web Server on port", config.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), router.GetRoutes()))
}

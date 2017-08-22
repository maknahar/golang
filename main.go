package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/maknahar/go-web-skeleton/internal/router"
)

func main() {
	fmt.Println("Hello!")
	log.Fatal(http.ListenAndServe(":9000", router.GetRoutes()))
}

package main

import (
	"log"
	"net/http"

	"gitlab.com/indie-developers/go-api-standard-library/api/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/healthcheck", &controllers.HealthCheck{})

	log.Println("app is running!")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("app has stopped, %s", err.Error())
	}
}

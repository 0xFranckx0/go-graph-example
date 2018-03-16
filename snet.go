package main

import (
	log "github.com/Sirupsen/logrus"
	server "github.com/coyotte-test/socnet"
	"net/http"
	"os"
)

//This is the main funtion that starts the API.
func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("starting API")

	router := server.NewRouter()
	var port string
	if port = os.Getenv("SNETPORT"); port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}

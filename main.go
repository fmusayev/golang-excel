package main

import (
	"flag"
	"net/http"

	"github.com/fmusayev/golang-excel/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	port   *string
	logFmt *string
)

func init() {
	godotenv.Load()
	port = flag.String("port", ":8080", "Port number to run")
	// Set log format to json if you are using tools for distributed logging (e.g. Datadog)
	logFmt = flag.String("log", "default", "Logging format: json")
}

func main() {
	flag.Parse()
	if *logFmt == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.Info("Application starting...")

	router := mux.NewRouter()
	handler.NewHandler(router)

	log.Info("Server listen at " + *port)
	log.Fatal(http.ListenAndServe(*port, router))
}

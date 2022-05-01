package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	hub := newHub(logrus.NewEntry(logger))
	go hub.Run()

	router := mux.NewRouter()

	router.Use(LoggingMiddleware(logger))
	router.HandleFunc("/chat", ClientHandler(hub))

	logger.Infof("Serving chat backend at %s", *addr)
	if err := http.ListenAndServe(*addr, router); err != nil {
		logger.WithError(err).Fatalln("starting server")
	}
}

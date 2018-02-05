package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	finish := make(chan bool)

	server8001 := http.NewServeMux()
	server8001.HandleFunc("/hello", hello)

	server8002 := http.NewServeMux()
	server8002.HandleFunc("/health", health)

	go func() {
		http.ListenAndServe(":8001", server8001)
	}()

	go func() {
		http.ListenAndServe(":8002", server8002)
	}()

	<-finish
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    "/hello",
		"method": "GET",
	}).Info("request made")
	w.Write([]byte("Hello World!"))
}

func health(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    "/healthy",
		"method": "GET",
	}).Info("request made")
	w.Write([]byte("I'm healthy!"))
}

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	finish := make(chan bool)

	server8001 := mux.NewRouter()
	server8001.HandleFunc("/hello", hello)
	server8001.NotFoundHandler = http.HandlerFunc(notFound)

	server8002 := mux.NewRouter()
	server8002.HandleFunc("/health", health)
	server8002.NotFoundHandler = http.HandlerFunc(notFound)

	go func() {
		log.Fatal(http.ListenAndServe(":8001", server8001))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8002", server8002))
	}()

	<-finish
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Error("404 not found")
	w.Write([]byte("404 Not Found"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Info("request made")
	w.Write([]byte("Hello World!"))
}

func health(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Info("request made")
	w.Write([]byte("I'm healthy!"))
}

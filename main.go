package main

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/joatmon08/hello/api"
	log "github.com/sirupsen/logrus"
)

var Version = "development"

func version(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Info("request made")
	w.Write([]byte(fmt.Sprintf("Version: %s", Version)))
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	finish := make(chan bool)

	server8001 := mux.NewRouter()
	server8001.HandleFunc("/", version)
	server8001.HandleFunc("/hello", api.Hello)
	server8001.HandleFunc("/phone", api.Phone)
	server8001.HandleFunc("/cpu", api.GenerateCPU)
	server8001.NotFoundHandler = http.HandlerFunc(api.NotFound)

	server8002 := mux.NewRouter()
	server8002.HandleFunc("/health", api.Health)
	server8002.NotFoundHandler = http.HandlerFunc(api.NotFound)

	go func() {
		log.Fatal(http.ListenAndServe(":8001", server8001))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8002", server8002))
	}()

	<-finish
}

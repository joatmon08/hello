package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	CPUTestDuration      = "5m"
	DefaultTargetService = "nginx"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	finish := make(chan bool)

	server8001 := mux.NewRouter()
	server8001.HandleFunc("/hello", Hello)
	server8001.HandleFunc("/phone", Phone)
	server8001.HandleFunc("/cpu", GenerateCPU)
	server8001.NotFoundHandler = http.HandlerFunc(NotFound)

	server8002 := mux.NewRouter()
	server8002.HandleFunc("/health", Health)
	server8002.NotFoundHandler = http.HandlerFunc(NotFound)

	go func() {
		log.Fatal(http.ListenAndServe(":8001", server8001))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8002", server8002))
	}()

	<-finish
}

func GenerateCPU(w http.ResponseWriter, r *http.Request) {
	var testTime time.Duration
	if r.URL.Query().Get("testTime") != "" {
		testTime, _ = time.ParseDuration(r.URL.Query().Get("testTime"))
	} else {
		testTime, _ = time.ParseDuration(CPUTestDuration)
	}

	log.WithFields(log.Fields{
		"uri":      r.URL.Path,
		"method":   r.Method,
		"host":     r.Host,
		"testTime": testTime,
	}).Info("request made")

	done := make(chan int)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}
	time.Sleep(testTime)
	close(done)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Error("404 not found")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("API Endpoint Not Found"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Info("request made")
	w.Write([]byte("Hello World!"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"uri":    r.URL.Path,
		"method": r.Method,
		"host":   r.Host,
	}).Info("request made")
	w.Write([]byte("I'm healthy!"))
}

func Phone(w http.ResponseWriter, r *http.Request) {
	var targetService string
	if r.URL.Query().Get("targetService") != "" {
		targetService = r.URL.Query().Get("targetService")
	} else {
		targetService = DefaultTargetService
	}

	log.WithFields(log.Fields{
		"uri":            r.URL.Path,
		"method":         r.Method,
		"host":           r.Host,
		"serviceToPhone": targetService,
	}).Info("request made")

	targetURL := fmt.Sprintf("http://%s", targetService)

	request, _ := http.NewRequest("GET", targetURL, bytes.NewBuffer(nil))
	client := &http.Client{}
	_, err := client.Do(request)
	if err != nil {
		log.WithFields(log.Fields{
			"uri":    r.URL.Path,
			"method": r.Method,
			"host":   r.Host,
			"err":    err,
		}).Errorf("request made to %s failed", targetURL)
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("I could not connect to %s!", targetURL)
		w.Write([]byte(message))
	} else {
		message := fmt.Sprintf("I connected to %s!", targetURL)
		w.Write([]byte(message))
	}
}

package main

import (
	"net/http"
)

func main() {
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
	w.Write([]byte("Hello World!"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm healthy!"))
}

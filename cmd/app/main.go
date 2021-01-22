package main

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	hello := http.HandlerFunc(helloServer)

	mux.Handle("/hello", mw(hello))

	log.Println("Listening on port 8080")

	http.ListenAndServe(":8080", mux)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("hi2", "hello2").Info("hello")
	w.Write([]byte("hello world!"))
}

func mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logrus.WithField("hi", "hello").Info("hello")

		next.ServeHTTP(w, r)
	})
}

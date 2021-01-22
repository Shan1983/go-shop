package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	hello := http.HandlerFunc(helloServer)

	mux.Handle("/hello", mw(hello))

	log.Println("Listening on port 8080")

	http.ListenAndServe(":8080", mux)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!!"))
}

func mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s %s", r.Host, r.Proto, r.URL)

		next.ServeHTTP(w, r)
	})
}

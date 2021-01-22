package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloServer)

	log.Println("Listening on port 8080")

	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
	fmt.Fprint(w, "Hello World!")
}

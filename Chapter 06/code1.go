package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Hello World")
}

func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

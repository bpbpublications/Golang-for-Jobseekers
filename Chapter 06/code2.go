package main

import (
	"fmt"
	"log"
	"net/http"
)

type HandleViaStruct struct{}

func (*HandleViaStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Hello World via Struct")
}

func main() {
	log.Print("Hello world sample started.")
	http.Handle("/", &HandleViaStruct{})
	http.ListenAndServe(":8080", nil)
}

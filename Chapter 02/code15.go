package main

import (
	"fmt"
)

var messages = make(chan string)

func main() {
	go createPing()

	msg := <-messages
	fmt.Println(msg)
}

func createPing() {
	messages <- "ping"
}

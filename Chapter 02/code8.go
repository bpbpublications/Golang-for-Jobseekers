package main

import "fmt"

func main() {
	l := LoveMessagePrinter{}
	PrintSomething(l)

	t := TroubleMessagePrinter{}
	PrintSomething(t)
}

func PrintSomething(m messagePrinter) {
	m.Print()
}

type messagePrinter interface {
	Print()
}

type LoveMessagePrinter struct{}

func (l LoveMessagePrinter) Print() {
	fmt.Println("I love Golang")
}

type TroubleMessagePrinter struct{}

func (t TroubleMessagePrinter) Print() {
	fmt.Println("I am still confused by this")
}

func (t TroubleMessagePrinter) AdditionalFunc() {
	fmt.Println("additionalFunc")
}

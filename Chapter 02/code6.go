package main

import "fmt"

// misc code snippets - for function declaration
// refer to subsection in chapter 2 about functions

func main() {
	exampleFunc()
	anotherExampleFunc()
	exampleFunc2("aa", 1)
	exampleFunc3(true, "aa", "bb", 1, 2)
	exampleFunc4(true, "aa", "bb", 1, 2)
	exampleFunc5("aa", "bb", "cc")
}

func exampleFunc() {
	fmt.Println("exampleFunc")
}

func anotherExampleFunc() {
	fmt.Println("anotherExampleFunc")
	a := 10
	if a == 10 {
		fmt.Println("Hello from func")
		return
	}
	fmt.Println("Bye")
}

func exampleFunc2(a string, b int) {
	fmt.Println(a)
	fmt.Println(b)
}

func exampleFunc3(a bool, b string, c string, d int, e int) {
	fmt.Println(a)
	fmt.Println(b)
}

func exampleFunc4(a bool, b, c string, d, e int) {
	fmt.Println(a)
	fmt.Println(b)
}

func exampleFunc5(a ...string) {
	fmt.Println(a)
}

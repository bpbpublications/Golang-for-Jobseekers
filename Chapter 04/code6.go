package main

// modded algo to store intermediate values to generate fibonacci seq

import "fmt"

var store = map[int]int{
	1: 1,
	2: 1,
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if store[n] != 0 {
		return store[n]
	}
	val := fibonacci(n-1) + fibonacci(n-2)
	store[n] = val
	return val
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21,...
func main() {
	fmt.Println(fibonacci(100))
}

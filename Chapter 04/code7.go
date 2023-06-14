package main

// another approach to generate fibonacci by tabulation (don't store all intermediate results)

import "fmt"

func fibonacciTabulate(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	previous1 := 1
	previous2 := 1
	currentVal := 0
	for i := 3; i <= n; i++ {
		currentVal = previous1 + previous2
		previous1 = previous2
		previous2 = currentVal
	}
	return currentVal
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21,...
func main() {
	fmt.Println(fibonacciTabulate(100))
}

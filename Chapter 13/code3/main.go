package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func FindBigger[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(FindBigger[int32](12, 15))
	fmt.Println(FindBigger[float32](1.5, 8.6))
	fmt.Println(FindBigger[string]("ab", "cc"))
}

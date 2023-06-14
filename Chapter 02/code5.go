package main

import "fmt"

// misc code snippets - for variable declaration
// refer to subsection about variables in chapter 2
//
// code snippets that reused names in earlier part of chapter might be renamed with numbers behind it
// e.g. primes -> primes2, primes3 etc
//
// additional println statements are added to ensure this golang file is "runnable"
// golang will not run if it detects variables declared in code but not used

func main() {
	// addition of type converted numbers
	var varInt64 int64 = 12
	varInt := 12
	summedInt := varInt + int(varInt64)
	fmt.Println(summedInt)

	// overflow
	varInt = 256
	transformed := uint8(varInt)
	fmt.Println(transformed)

	// declaring bool
	varBoolTrue := true
	varBoolFalse := false
	fmt.Println(varBoolTrue)
	fmt.Println(varBoolFalse)

	// declaring string
	varString := "TestVariable"
	fmt.Println(varString)

	// declaring fixed size array
	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
	// Printed value of primes variable:
	// [2 3 5 7 11 0]

	// declaring slice and adding a new number to it
	primes2 := []int{2, 3, 5, 7, 11}
	primes2 = append(primes2, 12)
	fmt.Println(primes2)

	// declaring slice and subsection-ing it
	primes3 := []int{2, 3, 5, 7, 11, 2}
	primesMiddle := primes3[2:4]
	primesFront := primes3[:3]
	primesBack := primes3[3:]
	fmt.Println(primesMiddle)
	fmt.Println(primesFront)
	fmt.Println(primesBack)
	// Printing of variables via fmt.Println
	// [5 7]
	// [2 3 5]
	// [7 11 2]

	// declaring/initializing 2d arrays
	var yoyo [][]int
	numbers := [][]int{{1, 2}, {3, 4}}
	fmt.Print(yoyo)
	fmt.Println(numbers)

	// different way of declaring and initializing slice
	primes4 := make([]int, 6)
	fmt.Println(primes4)
	primes4[0] = 2
	primes4[1] = 3
	fmt.Println(primes4)

	// initializing a golang map
	mappedItems := map[string]string{}
	mappedItems["test"] = "test"

	// alternative way to initialize golang map
	mappedItems2 := make(map[string]string)
	mappedItems2["test"] = "test"

	// other examples of declaring maps
	mappedItems3 := make(map[string]int)
	mappedItems4 := make(map[int]string)
	mappedItems5 := make(map[string]bool)
	mappedItems6 := make(map[string]func(a int) int)
	fmt.Println(mappedItems3)
	fmt.Println(mappedItems4)
	fmt.Println(mappedItems5)
	fmt.Println(mappedItems6)

}

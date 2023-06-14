package main

import "fmt"

func main() {
	items := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	for idx, x := range items {
		fmt.Println(idx)
		fmt.Println(x)
	}
}

package main

import (
	"fmt"

	temp "github.com/go-yaml/yaml"
)

func main() {
	fmt.Println("Hello, 世界")
	output, _ := temp.Marshal(map[string]string{"a": "b", "c": "d"})
	fmt.Println(string(output))
}

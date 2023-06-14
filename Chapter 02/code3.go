package main

import (
	"fmt"

	"github.com/go-yaml/yaml"
)

func main() {
	fmt.Println("Hello, 世界")
	output, _ := yaml.Marshal(map[string]string{"a": "b", "c": "d"})
	fmt.Println(string(output))
}

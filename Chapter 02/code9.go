package main

import "fmt"

func main() {
	allTasks := []Subtask{Subtask{Status: "incomplete"}, Subtask{Status: "completed"}}
	for idx, x := range allTasks {
		fmt.Println(idx)
		if x.Status != "completed" {
			fmt.Println("Main task is still incomplete")
		}
	}

}

type Subtask struct {
	Param1 string
	Param2 string
	Status string
}

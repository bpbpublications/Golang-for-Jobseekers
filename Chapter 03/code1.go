package main

// Contains code to manipulate singly linked lists

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}
	aa.Next = &bb
	bb.Next = &cc
}

func Print(root *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		fmt.Println(nodeWalk.Value)
		nodeWalk = nodeWalk.Next
	}
	fmt.Println(nodeWalk.Value)
}

func Len(root *Node) int {
	nodeWalk := root
	count := 1
	for nodeWalk.Next != nil {
		count = count + 1
		nodeWalk = nodeWalk.Next
	}
	return count
}

func Append(root *Node, newNode *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		nodeWalk = nodeWalk.Next
	}
	nodeWalk.Next = newNode
}

func Insert(root *Node, loc int, newNode *Node) error {
	nodeWalk := root
	counter := 0
	for nodeWalk.Next != nil {
		if counter == loc-1 {
			temp := nodeWalk.Next
			newNode.Next = temp
			nodeWalk.Next = newNode
			return nil
		}
		nodeWalk = nodeWalk.Next
		counter = counter + 1
	}
	return fmt.Errorf("Went past no of elements in list")
}

func Delete(root *Node, loc int) error {
	nodeWalk := root
	previousWalk := root
	counter := 0
	for nodeWalk != nil {
		if counter == loc {
			previousWalk.Next = nodeWalk.Next
		}
		counter = counter + 1
		previousWalk = nodeWalk
		nodeWalk = nodeWalk.Next
		return nil
	}
	return fmt.Errorf("Went past the expected list")
}

func Search(root *Node, val int) *Node {
	nodeWalk := root
	for nodeWalk.Next != nil {
		if nodeWalk.Value == val {
			return nodeWalk
		}
		nodeWalk = nodeWalk.Next
	}
	return nil
}

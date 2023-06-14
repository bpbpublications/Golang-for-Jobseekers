package main

// contains code to handle doubly linked lists

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	bb.Next = &cc
	bb.Previous = &aa
	cc.Previous = &bb

}

func Insert(root *Node, newNode *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node
	for n != nil {
		if counter == loc {
			newNode.Next = n
			newNode.Previous = p
			if n.Next != nil {
				n.Next.Previous = newNode
			}
			if p != nil {
				p.Next = newNode
				return root
			}
			return newNode
		}
		p = n
		n = n.Next
		counter = counter + 1
	}
	if counter == loc {
		newNode.Previous = p
		p.Next = newNode
		return root
	}
	return nil
}

func Delete(root *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node
	for n != nil {
		if counter == loc {
			if p == nil {
				temp := n.Next
				n.Next = nil
				n.Previous = nil
				return temp
			}
			p.Next = n.Next
			n.Next = nil
			n.Previous = nil
			return root
		}
		p = n
		n = n.Next
		counter = counter + 1
	}
	return nil
}

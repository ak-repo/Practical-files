package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

// append
func (ll *LinkedList) Append(value int) bool {
	newNode := &Node{Data: value}
	if ll.Head == nil {
		ll.Head = newNode

	} else {
		ll.Tail.Next = newNode
	}
	ll.Tail = newNode
	ll.Tail.Next = ll.Head

	return true

}

func (ll *LinkedList) Display() {
	curr := ll.Head

	for {
		fmt.Println("d: ", curr.Data)
		curr = curr.Next
		if curr == ll.Head {
			break
		}
	}
}

package main

import "fmt"

//linedlist

type Node struct {
	Data int
	Next *Node
}

type linkedlist struct {
	Head *Node
	Tail *Node
}

// appened
func (l *linkedlist) Append(value int) {

	newNode := &Node{Data: value}

	if l.Head == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
}

func (l *linkedlist) Display() {
	curr := l.Head

	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}

func (l *linkedlist) removeDuplicate() {

	curr := l.Head

	for curr.Next != nil {
		if curr.Data == curr.Next.Data {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

}



func main() {

	var list linkedlist

	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(3)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	// list.Display()
	list.Display()

}

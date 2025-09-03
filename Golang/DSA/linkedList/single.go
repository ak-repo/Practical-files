package main

// import "fmt"

// //Node

// type Node struct {
// 	Data int
// 	Next *Node
// }

// // linked list

// type LinkedList struct {
// 	Head   *Node
// 	Tail   *Node
// 	Length int
// }

// // method append
// func (lI *LinkedList) Append(data int) {

// 	newNode := &Node{Data: data}
// 	if lI.Head == nil {
// 		lI.Head = newNode

// 	} else {
// 		lI.Tail.Next = newNode
// 	}
// 	lI.Tail = newNode
// 	lI.Length++
// }

// // method display
// func (lI *LinkedList) Display() {
// 	if lI.Head == nil {
// 		fmt.Println("empty list")
// 		return
// 	}
// 	current := lI.Head
// 	fmt.Print("head ->")
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.Data)
// 		current = current.Next
// 	}
// 	fmt.Println("<-tail")

// }

// //method delete

// func (lI *LinkedList) Delete(data int) bool {
// 	current := lI.Head
// 	var prev *Node = nil

// 	if current.Data == data {
// 		lI.Head = current.Next
// 		// if there is one element only
// 		if lI.Head == nil {
// 			lI.Tail = nil
// 		}
// 		return true
// 	}

// 	for current != nil && current.Data != data {
// 		prev = current
// 		current = current.Next

// 	}

// 	if current == nil {
// 		return false // no value found

// 	}
// 	if current == lI.Tail {
// 		lI.Tail = prev
// 		lI.Tail.Next = nil
// 		return true
// 	}

// 	//if the case were value not head and tail
// 	prev.Next = current.Next
// 	return true

// }

// // method insert after a node

// func (lI *LinkedList) InsertAfter(nextTo, data int) bool {

// 	newNode := &Node{Data: data}
// 	current := lI.Head
// 	for current != nil && current.Data != nextTo {
// 		current = current.Next
// 	}
// 	fmt.Println("curr", current)

// 	if current == nil {
// 		return false
// 	}
// 	if current == lI.Tail {
// 		lI.Tail.Next = newNode
// 		lI.Tail = newNode
// 		return true
// 	}

// 	newNode.Next = current.Next
// 	current.Next = newNode
// 	return true

// }

package main

import "fmt"

//node
type DNode struct {
	Prev *DNode
	Data int
	Next *DNode
}

// double linkedlist

type DoubleList struct {
	Head *DNode
	Tail *DNode
}

// method append on tail
func (dl *DoubleList) Append(data int) {
	var newNode = &DNode{Data: data}

	if dl.Head == nil {
		dl.Head = newNode

	} else {
		dl.Tail.Next = newNode
		newNode.Prev = dl.Tail

	}
	dl.Tail = newNode
}

// method display forward
func (dl *DoubleList) DisplayForward() {

	current := dl.Head

	if current == nil {
		fmt.Println("empty list")
		return
	}
	fmt.Print("head ")
	for current != nil {
		fmt.Printf("%d-> ", current.Data)
		current = current.Next
	}
	fmt.Println(" tail")

}

// method display backword

func (dl *DoubleList) DisplayBackward() {

	current := dl.Tail
	if dl.Head == nil {
		fmt.Println("empty list")
		return
	}
	fmt.Print("Tail ")
	for current != nil && dl.Head.Next != nil {

		fmt.Printf("%d ->", current.Data)
		current = current.Prev

	}
	fmt.Println(" Head")

}

// method delete

func (dl *DoubleList) Delete(data int) bool {

	current := dl.Head

	// data is == head case
	if current.Data == data {
		dl.Head = current.Next
		dl.Head.Prev = nil

		if dl.Head == nil {
			dl.Tail = nil
		}
		return true
	}

	for current != nil && current.Data != data {
		current = current.Next
	}

	if current == nil {
		return false
	}

	// data is == tail case
	if current == dl.Tail {
		dl.Tail = current.Prev
		dl.Tail.Next = nil
		return true
	}

	//others
	current.Prev.Next = current.Next
	current = nil

	return true
}

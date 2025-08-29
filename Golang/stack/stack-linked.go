package main

import (
	"errors"
	"fmt"
)

// stack implemented using Linkedlist

type Node struct {
	Data int
	Next *Node
}

type Stack struct {
	Top *Node
}

//method diplay

func (s *Stack) Display() {

	current := s.Top

	if current == nil {
		fmt.Println("empty stack")
		return
	}

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}

}

// method Push() <-
func (s *Stack) Push(data int) {
	newNode := &Node{Data: data}

	if s.Top == nil {
		s.Top = newNode
	} else {
		newNode.Next = s.Top
		s.Top = newNode

	}
}

//method pop ->

func (s *Stack) Pop() (int, error) {

	if s.Top == nil {
		return 0, errors.New("stack underflow")
	}
	popped := s.Top.Data
	s.Top = s.Top.Next
	return popped, nil

}

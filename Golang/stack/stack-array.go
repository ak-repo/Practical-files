package main

import (
	"errors"
	"fmt"
)

// struct

type Stack struct {
	arr  []int
	top  int
	size int
}

// method init
func NewStack(size int) *Stack {

	return &Stack{
		arr:  make([]int, size),
		top:  -1,
		size: size,
	}
}

// method for Push elemenst
func (s *Stack) Push(value int) error {

	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.top++
	s.arr[s.top] = value
	return nil
}

// method for remove last element
func (s *Stack) Pop() (int, error) {

	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	}

	value := s.arr[s.top]
	s.top--
	return value, nil
}

// check stack is empty

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

//check stack is full

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

// method display stack top to bottom

func (s *Stack) Display() {

	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	fmt.Println("stack top to bottom")

	for i := s.top; i >= 0; i-- {
		fmt.Println(s.arr[i])
	}

}

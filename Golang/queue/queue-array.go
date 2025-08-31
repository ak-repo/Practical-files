package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr    []int
	front  int
	rear   int
	length int
	size   int
}

// new queue init function

func NewQueue(size int) *Queue {
	return &Queue{
		arr:    make([]int, size),
		front:  0,
		rear:   -1,
		size:   size,
		length: 0,
	}
}

//method for adding elemets

func (q *Queue) Enqueue(value int) {
	if q.IsFull() {
		fmt.Println("queue is full")
		return
	}

	q.rear = (q.rear + 1) % q.size
	q.arr[q.rear] = value
	q.length++

}

//method for removing elements from fornt

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}

	value := q.arr[q.front]
	q.arr[q.front] = 0

	q.front = (q.front + 1) % q.size //circulor increment
	q.length--
	return value, nil

}

//method for checking queue is full and empty

func (q *Queue) IsFull() bool {
	return q.size == q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

// method for display queue

func (q *Queue) Display() {
	if q.IsEmpty() {
		fmt.Println("empty Queue")
		return
	}

	for i := 0; i < q.length; i++ {
		index := (q.front + i) % q.size
		fmt.Println(q.arr[index])
	}
}

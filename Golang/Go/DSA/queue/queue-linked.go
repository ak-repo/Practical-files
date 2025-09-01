package main

// import "fmt"

// // queue using linkedlist
// type Node struct {
// 	Data int
// 	Next *Node
// }

// type Queue struct {
// 	Front  *Node
// 	Rear   *Node
// 	Length int
// }

// // method for adding elements in Rear

// func (q *Queue) Enqueue(data int) {

// 	newNode := &Node{Data: data}
// 	q.Length++
// 	// incase queue is empty
// 	if q.Front == nil {
// 		q.Front = newNode
// 		q.Rear = newNode
// 		return
// 	}

// 	q.Rear.Next = newNode
// 	q.Rear = newNode
// }

// // method for remove element form Front
// func (q *Queue) Dequeue() {
// 	if q.Front == nil {
// 		fmt.Println("Queue is empty")
// 		return
// 	}

// 	q.Front = q.Front.Next
// 	q.Length--

// 	//edge case
// 	if q.Front == nil {
// 		q.Rear = nil
// 	}
// }

// // display

// func (q *Queue) Display() {
// 	if q.Front == nil {
// 		fmt.Println("Queue is empty")
// 		return
// 	}

// 	fmt.Print("Front -> ")
// 	for current := q.Front; current != nil; current = current.Next {
// 		fmt.Printf("%d ", current.Data)
// 	}
// 	fmt.Println("<- Rear")
// }

// // method for showing Front element

// func (q *Queue) Peek() (int, bool) {
// 	if q.Front == nil {
// 		return 0, false
// 	}

// 	return q.Front.Data, true
// }

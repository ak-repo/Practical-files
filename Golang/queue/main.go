package main

func main() {
	// var qu Queue

	// qu.Enqueue(1)
	// qu.Enqueue(2)
	// qu.Enqueue(3)
	// qu.Enqueue(4)

	// qu.Dequeue()
	// qu.Display()

	// //accessing fornt value
	// v, ok := qu.Peek()
	// if ok {
	// 	fmt.Println("Front value: ", v)
	// } else {
	// 	fmt.Println("empty queue")
	// }

	// fmt.Println("len: ", qu.Length)

	//----------------------------------------------------

	qu := NewQueue(10)
	qu.Enqueue(1)
	qu.Enqueue(2)
	qu.Enqueue(3)
	qu.Enqueue(4)

	qu.Dequeue()
	qu.Dequeue()

	qu.Display()

}

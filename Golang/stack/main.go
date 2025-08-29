package main

import "fmt"

func main() {

	var stack Stack

	stack.Push(20)
	stack.Push(40)
	stack.Push(60)
	stack.Push(80)
	stack.Push(100)

	stack.Display()
	x, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("x:", x)
	stack.Display()


}

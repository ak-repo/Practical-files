package main

import "fmt"

func main() {

	y := []int{ 1, 2,5}
	x := []int {3,6,9,0,0,0}
	

	MergeArray(x, 3, y, 3)

	fmt.Println("x:", x)

}

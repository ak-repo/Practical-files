//  Creating a multiplication table with all the elements in the array. So
// if your array is [2, 3, 7, 8, 10], you first multiply every element by 2,
// then multiply every element by 3, then by 7, and so on.

package main

import "fmt"

func main() {

	x := []int{2, 3, 7, 8, 10}

	table := make([][]int, len(x))

	for i, outerValue := range x {
		table[i] = make([]int, len(x))

		for j, innerValue := range x {
			table[i][j] = outerValue * innerValue
		}

	}

	fmt.Println("ta: ", table)
}

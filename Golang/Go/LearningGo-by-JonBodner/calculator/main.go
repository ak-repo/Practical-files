package main

import (
	"fmt"
	"strconv"
)

func add(a, b int) int   { return a + b }
func sub(a, b int) int   { return a - b }
func multi(a, b int) int { return a * b }
func div(a, b int) int   { return a / b }

var calMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": multi,
	"/": div,
}

func main() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, exp := range expressions {

		if len(exp) != 3 {
			fmt.Println("not a valid calculation")
			continue
		}

		// values
		val1, ok := strconv.Atoi(exp[0])

		if ok != nil {
			fmt.Println(ok)
			continue
		}
		val2, ok := strconv.Atoi(exp[2])

		if ok != nil {
			fmt.Println(ok)
			continue
		}

		operator := exp[1]
		checkOp, err := calMap[operator]
		if !err {
			fmt.Println("Unsupported operation", err)
			continue

		}
		result := checkOp(val1, val2)

		fmt.Println(result)

	}

}

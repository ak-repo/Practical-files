package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		return
	}

	operator := os.Args[1]
	operand1, err1 := strconv.Atoi(os.Args[2])
	operand2, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid operands")
		return
	}

	result := 0
	switch operator {
	case "add":
		result = operand1 + operand2
	case "sub":
		result = operand1 - operand2
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("%d %s %d = %d\n", operand1, operator, operand2, result)
}

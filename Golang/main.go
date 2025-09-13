package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	x := sortArray([]int{5, 2, 3, 1})
	fmt.Println("x:", x)

}

// token generation
func GenerateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Token generate Failed: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)

}

func sortArray(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	pivote := nums[mid]

	// storing
	var less []int
	var greate []int

	for i, v := range nums {

		if i == mid {
			continue
		}

		if v > pivote {
			greate = append(greate, v)
		} else {
			less = append(less, v)
		}
	}

	return append(append(sortArray(less), pivote), sortArray(greate)...)

}

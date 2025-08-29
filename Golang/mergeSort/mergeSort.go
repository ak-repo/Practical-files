package main

import "fmt"

func MergeArray(arr1 []int, n1 int, arr2 []int, n2 int) {

	last := (n1+n2)-1

	for n1 > 0 && n2 > 0 {

		if arr1[n1-1] > arr2[n2-1] {
			arr1[last] = arr1[n1-1]
			n1--
		} else {
			arr1[last] = arr2[n2-1]
			n2--
		}
		last--
	}

	//if anything left over

	for n2 > 0 {
		arr1[last] = arr2[n2-1]
		last--
		n2--
	}

	fmt.Println("merg", arr1)

}

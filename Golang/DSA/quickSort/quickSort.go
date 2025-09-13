package main

var Count int

// case that choose 1st elemnt   O(n2)
// func QuickSort(arr []int) []int {

// 	Count++

// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	var less []int
// 	var greater []int
// 	pivot := arr[0]

// 	for _, v := range arr[1:] {

// 		if v > pivot {
// 			greater = append(greater, v)
// 		} else {
// 			less = append(less, v)
// 		}

// 	}

// 	return append(append(QuickSort(less), pivot), QuickSort(greater)...)

// }

// elemt take middle

func QuickSort(arr []int) []int {


	if len(arr) < 2 {
		return arr
	}
	var less []int
	var greater []int

	mid := len(arr) / 2
	pivot := arr[mid]

	for i, v := range arr {
		if i == mid {
			continue
		}

		if v > pivot {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)

}

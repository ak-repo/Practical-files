package main

import "fmt"

func main() {

	// single linkelist accessing

	// 	Input: l1 = [2,4,3], l2 = [5,6,4]
	// Output: [7,0,8]
	var sList LinkedList
	sList.Append(2)
	sList.Append(4)
	sList.Append(3)

	var sl LinkedList
	sl.Append(5)
	sl.Append(6)
	sl.Append(4)



	//----------------------------------------------------------------------------------------

	// var dList DoubleList
	// dList.Append(20)
	// dList.Append(40)
	// dList.Append(60)
	// dList.Append(80)

	// dList.DisplayForward()
	// dList.DisplayBackward()
	// dList.Delete(40)
	// dList.Delete(80)
	// dList.Delete(20)
	// dList.Delete(60)

	// dList.DisplayForward()

}

// remove sList duplicate elements,, it only work when same elements are in grouped together

func removeDuplicates(list *LinkedList) {

	var prev *Node = list.Head
	var current *Node = list.Head.Next
	if prev == nil || current == nil {
		return
	}

	for current != nil {

		if prev.Data == current.Data {
			prev.Next = current.Next
			current = prev.Next

		} else {
			prev = current
			current = current.Next
		}

	}

	//fix the tail pointer

	temp := list.Head
	for temp != nil && temp.Next != nil {
		temp = temp.Next
	}

	list.Tail = temp

}

// reversing linkedlist

func reverseLinkedlist(list *LinkedList) {
	

}


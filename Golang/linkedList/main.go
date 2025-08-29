package main

import "fmt"

func main() {

	// single linkelist accessing

	var sList LinkedList
	sList.Append(2)
	sList.Append(5)
	sList.Append(5)
	sList.Append(10)
	sList.Append(15)
	sList.Append(15)
	sList.Append(15)
	sList.Append(20)
	sList.Append(20)
	sList.Append(20)
	sList.Display()
	fmt.Println("len", sList.Length)

	// removeDupli(&sList)
	// sList.Display()

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

func removeDupli(list *LinkedList) {

	var prev *Node = list.Head
	var current *Node = list.Head.Next

	for current != nil {

		if prev.Data == current.Data {
			prev.Next = current.Next
			current = prev.Next

		} else {
			prev = current
			current = current.Next
		}

	}

	if list.Head.Data == list.Tail.Data {
		fmt.Println("true")
		list.Tail = nil
	}

}

package main

func main() {

	var ls LinkedList
	ls.Append(22)
	ls.Append(33)
	ls.Append(44)
	ls.Append(55)

	ls.Append(66)
	ls.Append(77)
	ls.Display()
}

// remove sList duplicate elements,, it only work when same elements are in grouped together

// func removeDuplicates(list *LinkedList) {

// 	var prev *Node = list.Head
// 	var current *Node = list.Head.Next
// 	if prev == nil || current == nil {
// 		return
// 	}

// 	for current != nil {

// 		if prev.Data == current.Data {
// 			prev.Next = current.Next
// 			current = prev.Next

// 		} else {
// 			prev = current
// 			current = current.Next
// 		}

// 	}

// 	//fix the tail pointer

// 	temp := list.Head
// 	for temp != nil && temp.Next != nil {
// 		temp = temp.Next
// 	}

// 	list.Tail = temp

// }

// // reversing linkedlist

// func reverseLinkedlist(list *LinkedList) {

// 	var prev *Node
// 	curr := list.Head

// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next

// 	}

// 	list.Head = prev

// }

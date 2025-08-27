package main

func main() {

	// single linkelist accessing

	var sList LinkedList
	sList.Append(2)
	sList.Append(5)
	sList.Append(12)
	sList.Append(9)
	sList.Display()

	sList.Delete(5)
	sList.Display()

	sList.InsertAfter(1000, 24)
	sList.Display()

}

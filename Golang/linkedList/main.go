package main

func main() {

	// single linkelist accessing

	// var sList LinkedList
	// sList.Append(2)
	// sList.Append(5)
	// sList.Append(12)
	// sList.Append(9)
	// sList.Display()

	// sList.Delete(5)
	// sList.Display()

	// sList.InsertAfter(1000, 24)
	// sList.Display()

	var dList DoubleList
	dList.Append(20)
	dList.Append(40)
	dList.Append(60)
	dList.Append(80)

	dList.DisplayForward()
	dList.DisplayBackward()
	dList.Delete(80)
	dList.DisplayForward()

}

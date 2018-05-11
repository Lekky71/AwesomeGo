package main

import "fmt"

func main(){
	testArrays()
}

func testArrays(){
	var nums [5]int
	nums[0] = 4
	nums[3] = 6
	fmt.Printf("The first element of the array is %d\n", nums[0])
	fmt.Printf("The last element of the array is %d\n", nums[len(nums)-1])

	//	note :: array length cannot be changed
	//	a := [3]int{1,2,3}
	//	b := [7]int{1,9,3,6} //assigns 0 to the remaining three elements
	//	c := [...]int{3,43,2,1,28} // ... make Go calculate the length itself

	//Defining double arrays
	//	doubleArray := [2][4]int{[4]int{3,45,0,3}, [4]int{2,4,6,1}} //first method
	//	dArray := [2][4]int{{2,4,5,3}, {4,5,3,2}} //easier method

	//A mutable array(ArrayList) is called Slice in Go
	//	var mySclice []int
	//	slicer := []byte {'a', 'b', 'c'}

	//Getting a slice from an array
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[6:9]
	fmt.Printf("%s\n", a)
	fmt.Printf("%s\n", b)

	c := ar[:6]
	d := ar[4:]
	fmt.Printf("%s\n", c)
	fmt.Printf("%s\n", d)
	//change a value in the slice
	c[2] = 'm' //this change affects the array a

	newArrslice := ar[:] //copy the entire array
	fmt.Printf("%s\n", newArrslice)
	//fmt.Printf("%s\n", b)

}

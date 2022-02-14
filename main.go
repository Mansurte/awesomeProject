// Package main for testing purposes . Ariel.
package main

import (
	"fmt"
	"strconv"
)

var (
	Px float64
	j  string
	k  string
	//a     bool
)

func arrays(phrase string) [3]string {
	testArray := [...]string{phrase, phrase, phrase}
	return testArray
}

func main() {
	// BASIC TYPES
	// Formatting floats to print as strings, concatenating strings, exponential notation
	Px = 367.4432
	j = strconv.FormatFloat(Px, 'E', -1, 32)
	k = "is the number of my house."
	k = "CAMBIO!"
	b := []byte(j + k)
	expo := 13e4
	fmt.Printf("The number of the expo is: %v\n", expo)
	fmt.Printf("%v, %T\n", j+k, j+k)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("The Price of the Future is $ %.2f so we need to buy\n", Px)
	fmt.Printf("The repeated phrase is:\n %v\n", arrays("Nooo!\n"))

	// ARRAYS
	// Tests with arrays
	var arrayTest = [...]string{"a", "b", "c"} // Can be [3] but it is better to use [...]
	fmt.Println(arrayTest)
	fmt.Printf("The second element is: %v\n", arrayTest[1])
	copyArrayTest := arrayTest
	copyArrayTest[1] = "b_copied"
	fmt.Printf("The second element of COPY is: %v\n", copyArrayTest[1])
	fmt.Printf("But the original remains!: %v\n", arrayTest[1])

	// Defining copy_array_test as a pointer
	pointerArrayTest := &arrayTest
	pointerArrayTest[1] = "b_pointed"
	fmt.Printf("The second element of pointed is now: %v\n", pointerArrayTest[1])
	fmt.Printf("And now the original changed!: %v\n", arrayTest[1])

	// SLICES
	// Same example as before but with slices... These always change the underlying when creating a "copy"
	var sliceTest = []string{"a", "b", "c"} // Note the empty brackets now, nor a number neither "..."
	copySliceTest := sliceTest
	copySliceTest[1] = "b_copied"
	fmt.Printf("The second element of COPY is: %v\n", copySliceTest[1])
	fmt.Printf("But the original remains!: %v\n", sliceTest[1])

	// Concatenating slices
	sliceTest = append(sliceTest, copySliceTest...) //  the "..." unpacks the slice into individual strings
	fmt.Printf("The concatenated slice is: %v\n", sliceTest)

	// To erase a middle element
	sliceWithoutSecondElement := append(sliceTest[:1], sliceTest[2:]...)
	fmt.Printf("The slice without the second element is %v\n", sliceWithoutSecondElement)
	//But... be careful with the change in the original array!!!
	fmt.Printf("The original slice array is now: %v.\nIts length is %v.\n", sliceTest, len(sliceTest))
	// The last element is still there in its original position!
}

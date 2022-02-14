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
	test_array := [...]string{phrase, phrase, phrase}
	return test_array
}

func main() {
	Px = 367.44321
	//a = true
	j = strconv.FormatFloat(Px, 'E', -1, 32)
	k = "is the number of my house."
	k = "CAMBIO!"
	b := []byte(j + k)
	expo := 13e4
	fmt.Printf("The number of the expo is: %v\n", expo)
	fmt.Printf("%v, %T\n", j+k, j+k)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("The Price of the Future is $ %.2f so we need to buy\n", Px)
	fmt.Printf("The phrase is:\n %v\n", arrays("Nooo!\n"))
	var array_test = [3]string{"a", "b", "c"}
	fmt.Println(array_test)
}

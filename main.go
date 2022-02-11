// Package main for testing purposes . Ariel.
package main

import (
	"fmt"
	"strconv"
)

var (
	inter int
	j     string
	k     string
	//a     bool
)

func main() {
	inter = 67
	//a = true
	j = strconv.Itoa(inter)
	k = "is the number of my house."
	b := []byte(j + k)
	expo := 13e4
	fmt.Printf("The number of the expo is: %v\n", expo)
	fmt.Printf("%v, %T\n", b, b)
}

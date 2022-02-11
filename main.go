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
	fmt.Printf("%v ", j)
	fmt.Printf("%v\n", k)
}

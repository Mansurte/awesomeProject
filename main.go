// Package main for testing purposes . Ariel.
package main

import (
	"fmt"
	"strconv"
)

var (
	inter int
)

func main() {
	inter = 67
	var a bool
	a = true
	var j string
	//var s2 string
	j = strconv.Itoa(inter)
	//j := 40
	fmt.Printf("%v, %T \n", j, j)
	fmt.Printf("%v, %T\n", a, a)
}

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
	var j string
	j = strconv.Itoa(inter)
	//j := 40
	fmt.Printf("%v, %T \n", j, j)
}

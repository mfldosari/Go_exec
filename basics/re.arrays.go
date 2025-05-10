package main

import (
	"fmt"
)

var arry1 = []int{}

func main() {
	myslice1 := arry1[:]
	myslice1 = append(myslice1, 1, 2)
	fmt.Println(myslice1[:])
}

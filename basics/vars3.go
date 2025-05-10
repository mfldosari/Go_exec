package main

import (
	"fmt"
)

// if type is not spicifed we can assign any type to the vars
func main() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

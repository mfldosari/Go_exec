package main

import (
	"fmt"
)

func main() {
	var this1, this2, this3 = "hello from this1", 56, true

	this4, this5, this6 := "hello from this4", 902, 78.35

	fmt.Println(this1)
	fmt.Println(this2)
	fmt.Println(this3)
	fmt.Println(this4)
	fmt.Println(this5)
	fmt.Println(this6)
}

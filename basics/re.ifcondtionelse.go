package main

import "fmt"

func main() {
	num1 := 89
	num2 := 90
	num3 := 90

	if num1 != num2 && (num3 == num2-num1) {
		fmt.Println("Hello")
	} else if num2 == num3 {
		fmt.Println("hey2")
	} else {
		fmt.Println("hey")

	}
}

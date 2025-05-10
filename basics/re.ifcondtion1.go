package main

import "fmt"

func main() {
	num1 := 89
	num2 := 90
	num3 := 1

	if num1 != num2 && (num3 == num2-num1) {
		fmt.Println("Hello")
	}
}

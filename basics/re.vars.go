package main

import (
	"fmt"
)

// we can define a var in three ways

// one using var <name of var> <type> = <value>
// two using var <name of var> = <value>
// three using <name of value> := <value>
func main() {
	name := "this is the name value"
	var x = 23
	var float float64 = 65.34

	fmt.Println(len(name) * 4)
	fmt.Println(x)
	fmt.Println(float * 2)

}

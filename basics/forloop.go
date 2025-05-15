package main

import (
	"fmt"
)

func loop(s string) []string {
	var mylist = []string{}
	for i := 0; i < len(s); i++ {
		mylist = append(mylist, string(s[i]))

	}
	return mylist
}

func main() {
	mylist := loop("this the test")
	first_string := mylist[0:4]
	fmt.Println(mylist)
	fmt.Println(first_string)

}

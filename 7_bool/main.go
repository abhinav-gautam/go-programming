package main

import (
	"fmt"
)

var x bool
func main() {

	fmt.Println(x) // zero value is false
	x=true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b)
}

package main

import (
	"fmt"
	"runtime"
)

var x int8

func main() {
	a := 45
	b := 5.465
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Println(b)
	fmt.Printf("%T\n",b)

	x = 127
	fmt.Println(x)

	// Runtime Package
	fmt.Println(runtime.GOARCH) // Architecture
	fmt.Println(runtime.GOOS) // Operating System

}


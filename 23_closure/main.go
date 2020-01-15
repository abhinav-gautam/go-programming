package main

import "fmt"

var x int	//Package level scope
func main() {
	y:=24 // Scope only in main
	y++
	x++
	fmt.Println(x,y)
	{
		z:=89 // Scope only in this code block
		z++
		fmt.Println(x,y,z)
	}
	// fmt.Println(z) // Error

	a := incrementer()
	b := incrementer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer() func() int{
	var x int // It has different scope for different variables
	return func() int {
		x++
		return x
	}
}
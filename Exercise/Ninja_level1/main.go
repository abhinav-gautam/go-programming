package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

type chickenBiryani int
var p chickenBiryani
var q int
func main() {

	// Exercise 1

	fmt.Println("Exercise 1")
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x,y,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Exercise 2

	fmt.Println("Exercise 2")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) // Zero Values

	//Exercise 3

	fmt.Println("Exercise 3")
	s := fmt.Sprintf("%v\t%v\t%v",a,b,c)
	fmt.Println(s)

	// Exercise 4

	fmt.Println("Exercise 4")
	fmt.Printf("%v\n%T\n",p,p)
	p = 42
	fmt.Println(p)

	//Exercise 5

	fmt.Println("Exercise 6")
	q = int(p)
	fmt.Printf("%v\n%T\n",q,q)

}

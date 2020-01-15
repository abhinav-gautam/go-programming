package main

import "fmt"

var a int
type abhinav int
var b abhinav // variable of type abhinav
func main() {
	a = 45
	b = 65
	fmt.Println("a =",a)
	fmt.Printf("%T\n",a)
	fmt.Println("b =",b)
	fmt.Printf("%T\n",b)

	// a = b // this will give error due to incompatible types

	a = int(b) // this is known as type conversion. term type casting is not used in go
	fmt.Println("a =",a)
	fmt.Printf("%T\n",a)
}

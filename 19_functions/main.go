package main

import (
	"fmt"
	"math"
)

func main() {
	foo()
	woo("Abhinav")

	defer woo("Gaurav") // this will run at the last

	fmt.Println(goo("Ansh"))
	x, y := hoo("Jayant", 20)
	fmt.Println(x, y)

	sum2 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum2)

	xi := []int{4, 9, 5, 2, 3, 8}
	fmt.Println(sum(xi...)) // Unfurling a slice

	// Anonymous Func
	func(){
		fmt.Println("This is anonymous function")
	}()

	func(x int){
		fmt.Println("Square of number:",math.Pow(float64(x),2))
	}(4)

	// Func Expression
	f := func(x float64) {
		fmt.Println("Sqrt of number:",math.Sqrt(x))
	}

	f(36)

	fmt.Printf("%T\n",f) // Func has type just like others

	// Returning a function
	k:= bar()
	fmt.Printf("%T\n",k)
	fmt.Println(bar()())

}

//func (receiver) identifier(parameters) (returns) { code }

func foo() {
	fmt.Println("Hello from foo")
}

func woo(s string) {
	fmt.Println("Hello", s, "from woo")
}

func goo(s string) string {
	return fmt.Sprintln("Hello", s, "from goo")
}

func hoo(s1 string, age int) (string, bool) {
	a := fmt.Sprintln("Hello", s1, "from hoo age is", age)
	return a, true
}

func sum(x ...int) int { // Variadic Parameter
	fmt.Printf("x=%v \t type=%T\n", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// Returning a function
func bar() func() int{
	return func() int {
		return 1997
	}
}
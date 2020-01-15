package main

import (
	"fmt"
	"math"
)

// Exercise 4
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old.")
}

// Exercise 5
type circle struct {
	radius float64
}
type square struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("Area of circle is:", s.area())
	case square:
		fmt.Println("Area of square is:", s.area())
	}
}

func main() {
	// Exercise 1
	fmt.Println(foo1())
	fmt.Println(bar1())

	// Exercise 2
	fmt.Println(foo2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
	fmt.Println(bar2([]int{10, 11, 12, 13, 41, 6, 5}))

	// Exercise 3
	defer bye()

	// Exercise 4
	p1 := person{
		first: "abhinav",
		last:  "gtm",
		age:   25,
	}
	p1.speak()

	// Exercise 5
	c1 := circle{
		radius: 5,
	}
	s1 := square{
		length: 10,
		width:  10,
	}
	info(c1)
	info(s1)

	// Exercise 6
	func() {
		fmt.Println("We are Anonymous!")
	}()

	// Exercise 7
	f := func() {
		fmt.Println("Function assigned to variable")
	}

	f()

	// Exercise 8
	g := it()
	g()

	// Exercise 9
	fmt.Println(incrementAdder(add, 8, 9))
}

// Exercise 1
func foo1() int {
	return 1997
}
func bar1() (int, string) {
	return 2019, "Abhinav"
}

// Exercise 2
func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar2(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// Exercise 3
func bye() {
	fmt.Println("Program: bye bye")
}

// Exercise 8
func it() func() {
	return func() {
		fmt.Println("Function returned by function")
	}
}

// Exercise 9
func add(x, y int) int {
	return x + y
}
func incrementAdder(f func(x, y int) int, x, y int) int {
	x++
	y++
	return f(x, y)
}

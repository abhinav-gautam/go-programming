package main

import (
	"fmt"
)

const (
	_  = iota
	kb1 = 1 << (iota * 10)
	mb1 = 1 << (iota * 10)
	gb1 = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1 // Bitshift operator
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)
}

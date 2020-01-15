package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d
	e
	f // Auto incrementing
)
const (
	g = iota
	h
	i
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

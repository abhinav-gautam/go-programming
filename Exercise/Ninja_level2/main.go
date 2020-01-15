package main

import "fmt"

// Exercise 3
const (
	k int = 45        // typed constant
	p     = "abhinav" // untyped constant
)
const (
	g = 2019 + iota
	h = 2019 + iota
	i = 2019 + iota
	j = 2019 + iota
)

func main() {

	// Exercise 1

	x := 45
	fmt.Printf("Decimal %d\nBinary %b\nHex %#x\n", x, x, x)

	// Exercise 2

	a := (56 == 56)
	b := (12 < 56)
	c := (45 > 85)
	d := (95 >= 15)
	e := (65 <= 45)
	f := (45 != 23)
	fmt.Println(a, b, c, d, e, f)

	// Exercise 4

	r := 52
	fmt.Printf("Decimal %d\nBinary %b\nHex %#x\n", r, r, r)
	q := r << 1 // Bitshift operator
	fmt.Printf("Decimal %d\nBinary %b\nHex %#x\n", q, q, q)

	// Exercise 5

	l := `I am raw string
I'm just like this`
	fmt.Println(l)

	// Exercise 6
	fmt.Println(g, h, i, j)
}

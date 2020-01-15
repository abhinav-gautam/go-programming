package main

import "fmt"

const x int = 45 // Typed constant
const y = "Untyped constant"

const (
	z = 42
	a = "blabla"
)

func main() {
	fmt.Println(x, y, z, a)
}

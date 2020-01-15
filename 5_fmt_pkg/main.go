package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // type
	fmt.Printf("%b\n", y) // binary
	fmt.Printf("%x\n", y) // hex
	fmt.Printf("%#x\n", y) // hex with 0x
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // as string
	fmt.Println(s)
	fmt.Printf("%T\n",s)
	fmt.Printf("%v", y) // default format
}

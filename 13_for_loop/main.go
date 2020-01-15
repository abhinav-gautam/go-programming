package main

import "fmt"

func main() {

	// for init; condition; post {}

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// a way
	a := 1
	for a < 10 {
		fmt.Println(a)
		a++
	}

	// other way
	b := 0
	for {
		b++
		if b > 100 {
			break
		}
		if b%2 != 0 {
			continue
		}
		fmt.Println(b)

	}

	//Nested Loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)

		}
	}

	// Printing ascii
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}

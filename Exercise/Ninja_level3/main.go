package main

import "fmt"

func main() {

	// Exercise 1

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

	// Exercise 2
	for i := 65; i <= 122; i++ {
		fmt.Printf("%v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}

	}

	// Exercise 3
	age := 1997
	for age <= 2019 {
		fmt.Println(age)
		age++
	}

	// Exercise 4
	age = 1997
	for {
		if age > 2019 {
			break
		}
		fmt.Println(age)
		age++
	}

	// Exercise 5
	n := 10
	for n <= 100 {
		fmt.Println("Remainder is", n%4)
		n++
	}

}

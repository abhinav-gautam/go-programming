package main

import "fmt"

func main() {
	//// Doesnt work
	//c:= make(chan int)
	//c <- 42
	//
	//fmt.Println(<-c)

	//// Unbuffered Channel Will work
	c := make(chan int)

	go func() {
		c <- 45
	}()

	fmt.Println(<-c)

	//// Buffered Channel Will work
	b := make(chan int, 2)

	b <- 56
	b <- 98
	fmt.Println(<-b)
	fmt.Println(<-b)

	//// Buffered Channel doesnt work
	d := make(chan int, 1)

	d <- 65
	fmt.Println(<-d)
	d <- 32
	d <- 26
	fmt.Println(<-d)

}

package main

import "fmt"

func main() {
	cb := make(chan int, 1)
	cs := make(chan<- int, 1) // Send only channel
	cr := make(<-chan int)    // Receive only channel
	cs <- 45

	//cr <- 89	// Error
	//fmt.Println(<-cs) // Error
	//fmt.Println(<-cr)

	fmt.Printf("cb %T\n", cb)
	fmt.Printf("cs %T\n", cs)
	fmt.Printf("cr %T\n", cr)

	//// Specific to general assignment - Error
	//cb = cr
	//cb = cs

	//// General to specific assignment - No error
	cr = cb
	cs = cb

	//// Specific to general conversion - Error
	//fmt.Printf("%T\n", (chan int)(cr))
	//fmt.Printf("%T\n", (chan int)(cs))

	//// General to specific conversion - No error
	fmt.Printf("%T\n", (chan<- int)(cb))
	fmt.Printf("%T\n", (<-chan int)(cb))
}

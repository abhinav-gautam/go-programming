package main

import "fmt"

func main() {
	c:=make(chan int)

	//send
	go send(c)

	//receive
	receive(c)

	fmt.Println("Program Ends")
}

func send(c chan <- int)  {
	c <- 45
}

func receive(c <- chan int){
	fmt.Println(<-c)
}
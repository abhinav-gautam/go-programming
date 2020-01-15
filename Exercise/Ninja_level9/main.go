package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Exercise 2
type person struct{
	firstName string
	age int
}

func (p *person) speak()  {
	fmt.Println("My name is",p.firstName,"and I am",p.age,"years old.")
}

type human interface {
	speak()
}

func saySomething(h human)  {
	h.speak()
}

func main() {

	wg.Add(2)
	// Exercise 1
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("GO1",i)
		}
		wg.Done()
	}()

	go foo()

	wg.Wait()

	// Exercise 2
	p1 := person{
		firstName:"Abhinav",
		age:22,
	}
	saySomething(&p1)
	// saySomething(p1) // Error as speak has pointer receiver

}


//Exercise 1
func foo(){
	for i := 0; i < 10; i++ {
		fmt.Println("GO2",i)
	}
	wg.Done()
}
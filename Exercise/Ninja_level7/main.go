package main

import "fmt"

type person struct {
	firstName string
	age int
}
func main() {
	p1 := person{
		firstName:"abhinav",
		age:25,
	}
	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}

func changeMe(p *person)  {
	(*p).firstName = "Ansh"
	(*p).age = 20
}

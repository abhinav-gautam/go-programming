package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

type secretAgent struct{
	person
	ltk bool
}

func (s secretAgent) speak()  {
	fmt.Println("I am",s.firstName,s.lastName)
}

func main() {
	sa1 := secretAgent{
		person:person{
			firstName:"Abhinav",
			lastName:"Gtm",
		},
		ltk:true,
	}

	p1 := person{
		firstName:"Ansh",
		lastName:"Bhawnani",
	}
	sa2 := secretAgent{
		person:p1,
		ltk:false,
	}
	sa1.speak()
	sa2.speak()
}

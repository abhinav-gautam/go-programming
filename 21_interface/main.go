package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am secretAgent of name", s.firstName, s.lastName)
}

func (p person) speak() {
	fmt.Println("I am person of name", p.firstName, p.lastName)
}

type human interface { // This is interface
	speak() // types associated with func speak are also of type human
}

func bar(h human) { // Polymorphism of func bar it accepts both person and secretAgent types as they're of type human
	switch h.(type) {
	case person:
		fmt.Println("Human(Person)", h.(person).firstName, "at bar") // Assertion of type person
	case secretAgent:
		fmt.Println("Human(secretAgent)", h.(secretAgent).firstName, "at bar")
	}
}
func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "Abhinav",
			lastName:  "Gtm",
		},
		ltk: true,
	}

	p1 := person{
		firstName: "Gaurav",
		lastName:  "Roy",
	}

	sa2 := secretAgent{
		person: person{
			firstName: "Ansh",
			lastName:  "Bhawnani",
		},
		ltk: false,
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(p1)
}

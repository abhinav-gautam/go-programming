package main

import "fmt"

type person struct {
	first string
	last string
	age int
}
type secretAgent struct {
	person
	ltk bool //License to kill
	kd float64 // Kill to death ratio
}
func main() {
	p1:= person{
		first:`Abhinav`,
		last:`Gtm`,
		age:24,
	}
	sa1:= secretAgent{
		person:person{
			first : "James",
			last:"Bond",
			age:34,
		},
		ltk:true,
		kd:0,
	}
	fmt.Println(p1)
	fmt.Println(p1.first,p1.last,p1.age)
	fmt.Println(sa1)
	fmt.Println(sa1.first,sa1.last,sa1.age,sa1.ltk,sa1.kd)

	//Anonymous Struct
	p2 := struct{
		firstName string
		age int
	}{
		firstName:"Kahlil",
		age:39,
	}
	fmt.Println(p2)
}

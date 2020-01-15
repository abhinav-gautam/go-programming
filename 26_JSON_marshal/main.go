package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName string
	Age int

}

func main() {
	p1:= person{
		FirstName:"Abhinav",
		LastName:"Gtm",
		Age:23,
	}

	p2 := person{
		FirstName:"Ansh",
		LastName:"Bhawnani",
		Age:22,
	}

	people:=[]person{p1,p2}
	bs,err := json.Marshal(people)
	if err!= nil{
		fmt.Println("Error",err)
	}
	fmt.Println(string(bs))
}

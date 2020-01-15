package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Age       int    `json:"Age"`
}

func main() {
	s:= `[{"FirstName":"Abhinav","LastName":"Gtm","Age":23},{"FirstName":"Ansh","LastName":"Bhawnani","Age":22}]`
	bs := []byte(s)

	//fmt.Printf("%T\n%T\n",s,bs)

	var people []person

	err := json.Unmarshal(bs,&people)

	if err!=nil{
		fmt.Println("Shit happened:",err)
	}

	fmt.Println(people)

	for _,v := range people {
		fmt.Println(v.FirstName,v.LastName,v.Age)
	}
}

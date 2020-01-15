package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass:="123456789"
	bs,err := bcrypt.GenerateFromPassword([]byte(pass),15)

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println("Hashed Password:",bs)

	givenPass := "12345678"
	err = bcrypt.CompareHashAndPassword(bs,[]byte(givenPass))
	if err!=nil{
		fmt.Println("Incorrect Password")
		return
	}
	fmt.Println("Successfully Logged in")
}

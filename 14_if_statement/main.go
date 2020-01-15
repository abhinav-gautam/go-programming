package main

import "fmt"

func main() {
	if x := 44; x==41{ // initialization
		fmt.Println("Bamm!!!")
	}else if x==42 || x==40{
		fmt.Println("Near miss")
	}else{
		fmt.Println("Not even close")
	}
	//fmt.Println(x) // wont work as scope of x is limited
}

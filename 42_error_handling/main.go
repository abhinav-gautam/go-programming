package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//// 1
	//n,err := fmt.Println("Start of program")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(n)

	//// 2
	//var ans1, ans2, ans3 string
	//
	//fmt.Println("Enter Name")
	//_,err := fmt.Scan(&ans1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Enter Age")
	//_,err = fmt.Scan(&ans2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Enter Fav Food")
	//_,err = fmt.Scan(&ans3)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ans1,ans2,ans3)

	//// 3
	f, err:= os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r:=strings.NewReader("Abhinav")
	io.Copy(f,r)
}

package main

import "fmt"

func main() {
	//a:= 42
	//
	//fmt.Println(a)
	//fmt.Println(&a) // & gives the address of memory where value is stored
	//
	//fmt.Printf("%T\n",a)
	//fmt.Printf("%T\n",&a) // its of pointer type - pointing to int
	//
	//b:= &a
	//
	//fmt.Println(b)
	//fmt.Println(&b)
	//
	//fmt.Printf("%T\n",b)
	//fmt.Printf("%T\n",&b)
	//
	//fmt.Println(*b) // * gives the value stored at an address
	//fmt.Println(*&b)
	//
	//*b = 49
	//fmt.Println(a) // a and b points to a same address

	////Without Pointers
	x:=42
	//fmt.Println("x before",x)
	//foo1(x)
	//fmt.Println("x  after",x)

	////With Pointers
	x=42
	fmt.Println("x before",x)
	foo2(&x)
	fmt.Println("x  after",x)


}

//func foo1(y int)  {
//	fmt.Println("y before",y)
//	y = 49
//	fmt.Println("y  after",y)
//}

func foo2(y *int)  {
	fmt.Println("y before",*y)
	*y = 49
	fmt.Println("y  after",*y)
}
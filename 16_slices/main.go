package main

import "fmt"

func main() {
	// x:= type{values} // Composite Literals

	x := []int{1, 32, 23, 43, 458, 6}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[2])

	// Looping the slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	// Slicing a slice
	fmt.Println(x[:])
	fmt.Println(x[1:5])

	//Appending a slice
	x = append(x, 18, 96, 12)
	fmt.Println(x)
	y := []int{798, 654, 321}
	x = append(x, y...)
	fmt.Println(x)

	// Deleting from slice
	x = append(x[:5], x[9:]...)
	fmt.Println(x)

	// Slice using make
	x = make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Multi dimensional array
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Printf("%v\n%v",xp[0],xp[1])
}

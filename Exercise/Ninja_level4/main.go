package main

import "fmt"

func main() {

	// Exercise 1
	arr := [5]int{}
	arr[0] = 1
	arr[1] = 5
	arr[2] = 4
	arr[3] = 3
	arr[4] = 9

	for i,v:= range arr{
		fmt.Println(i,v)
	}
	fmt.Printf("Type of array %T\n",arr)

	// Exercise 2
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i,v:= range xi{
		fmt.Println(i,v)
	}
	fmt.Printf("Type of Slice %T\n",xi)

	//Exercise 3
	fmt.Println(xi[:5])
	fmt.Println(xi[5:10])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])

	// Exercise 4
	xi = append(xi, 52)
	fmt.Println(xi)
	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)
	y := []int{56, 57, 58, 59, 60}
	xi = append(xi, y...)
	fmt.Println(xi)

	// Exercise 5
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(x[:3],x[6:]...)
	fmt.Println(y)

	//Exercise 6
	am := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		           ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`,
		           ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`,
		           ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`,
		           ` New Jersey`,` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`,
		           ` Oregon`,` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`,
		           ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println("Length",len(am))
	fmt.Println("Capacity",cap(am))

	for i := 0; i<len(am); i++ {
		fmt.Println(i,am[i])
	}

	// Exercise 7
	nm := [][]string{{"James", "Bond", "Shaken, not stirred"},{"Miss", "Moneypenny", "Helloooooo, James."}}

	for _,v:=range nm{
		for _,c:= range v{
			fmt.Printf("%v\t",c)
		}
		fmt.Println()
	}

	// Exercise 8
	mp := map[string][]string{
		`bond_james`:{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k,v := range mp{
		fmt.Printf("%v\n",k)
		for i,vs:=range v{
			fmt.Printf("\t\t%v\t%v\n",i,vs)
		}
	}

	// Exercise 9
	mp[`gtm_abhin`] = []string{`computer science`,`laptop`,`food`}
	for k,v := range mp{
		fmt.Printf("%v\n",k)
		for i,vs:=range v{
			fmt.Printf("\t\t%v\t%v\n",i,vs)
		}
	}

	// Exercise 10
	delete(mp,`no_dr`)
	for k,v := range mp{
		fmt.Printf("%v\n",k)
		for i,vs:=range v{
			fmt.Printf("\t\t%v\t%v\n",i,vs)
		}
	}
}

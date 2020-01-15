package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	// Exercise 1
	p1 := person{
		firstName: "Abhinav",
		lastName:  "gtm",
		favIceCream: []string{
			"Chocolate",
			"Double Chocolate",
			"Oreo",
		},
	}

	p2 := person{
		firstName: "Ansh",
		lastName:  "Bhawnani",
		favIceCream: []string{
			"Butterscotch",
			"American Nut",
			"Mega Nut",
		},
	}

	fmt.Println(p1.firstName, "Favorite Ice Cream")
	for _, v := range p1.favIceCream {
		fmt.Printf("%v\n", v)
	}
	fmt.Println(p2.firstName, "Favorite Ice Cream")
	for _, v := range p2.favIceCream {
		fmt.Printf("%v\n", v)
	}

	// Exercise 2
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	for _, v := range m {
		fmt.Println(v.firstName)
		for _, v2 := range v.favIceCream {
			fmt.Println("\t", v2)
		}
	}

	// Exercise 3
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(truck1, sedan1)

	fmt.Printf("truck1 \n doors:%v \n color:%v \n fourWheel:%v \n", truck1.doors, truck1.color, truck1.fourWheel)
	fmt.Printf("sedan1 \n doors:%v \n color:%v \n luxury:%v \n", sedan1.doors, sedan1.color, sedan1.luxury)

	// Exercise 4
	hacker1 := struct {
		codeName string
		isAnon   bool
		ability  []string
	}{
		codeName: "Rayark",
		isAnon:   true,
		ability: []string{
			"xss",
			"csrf",
			"sqli",
			"metasploit",
		},
	}
	fmt.Println(hacker1)
}

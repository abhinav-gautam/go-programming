package main

import "fmt"

func main() {
	m := map[string]int{ // Composite literal // map[keyType]valueType{}
		"Abhinav": 22,
		"Ansh":    21,
		"Jayant":  20,
	}

	fmt.Println(m["Abhinav"])
	fmt.Println(m["Ansh"])
	fmt.Println(m["Gaurav"])

	// Checking a value if it exist
	v, ok := m["Gaurav"]
	fmt.Println(v, ok)

	if v, ok := m["Ansh"]; ok {
		fmt.Println("Person Exist, Age:", v)
	} else {
		fmt.Println("Person doesnt exist")
	}

	// Adding a key-value to map
	m["Gaurav"] = 20

	//Printing key values
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Deleting from map
	delete(m,"Ansh")
	delete(m,"Shivani")
	fmt.Println(m)

	if v,ok=m["Shivani"]; ok{
		delete(m,"Shivani")
		fmt.Println("Deleted Successfully")
	}else{
		fmt.Println("Person Doesn't Exist")
	}
}

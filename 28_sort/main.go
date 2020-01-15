package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age int
}

type ByAge []Person
type ByName []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("Sorted Ints",xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println("Sorted Strings",xs)

	//// Custom Sort
	//// By age
	p1 := Person{"Abhinav",22}
	p2 := Person{"Ansh",20}
	p3 := Person{"Jayant",19}
	people:= []Person{p1,p3,p2}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	//// By name
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

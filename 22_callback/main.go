package main

import "fmt"

// Even sum

func main() {
	fmt.Println(evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func evenSum(f func(x ...int) int, vi ...int) int { // Callback means passing func as argument in a func
	var xi []int
	for _, v := range vi {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

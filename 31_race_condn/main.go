package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	fmt.Println("CPUs:",runtime.NumCPU())
	fmt.Println("Goroutines:",runtime.NumGoroutine())

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)
}

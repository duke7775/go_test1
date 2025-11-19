package main

import (
	"fmt"
	"sync"
)

func print_number(index int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("index:", index, "i:", i)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go print_number(int64(i), &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

package main

import (
	"fmt"
	"sync"
)

func print_number(index int64, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	for i := 0; i < 10; i++ {
		fmt.Println("index:", index, "i:", i)

	}
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go print_number(int64(i), &wg, &mu)
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}

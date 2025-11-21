package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func print_number(index int64) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < 10; i++ {
		fmt.Println("index:", index, "i:", i)

	}

}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go print_number(int64(i))
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	defer fmt.Println("第三个defer")
}

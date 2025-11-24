package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type channel struct {
	value int
}

func print_number(ch chan channel) {
	defer wg.Done()
	c := <-ch
	fmt.Println(c.value)
}

func main() {
	ch := make(chan channel, 5)
	for i := 0; i < 5; i++ {
		ch <- channel{value: i}
	}
	close(ch)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go print_number(ch)
	}
	wg.Wait()
	fmt.Println("Finished")
}

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
	wg.Add(5)
	for i := 0; i < 5; i++ {

		go print_number(ch)
	}
	close(ch)
	wg.Wait()
	fmt.Println("Finished")
}

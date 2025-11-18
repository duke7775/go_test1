package main

import (
	"fmt"
	"time"
)

var g_index int = 0

func print_number(index int64) {
	for i := 0; i < 10; i++ {
		fmt.Println("index:", index, "i:", i)
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go print_number(int64(g_index))
		g_index++
	}
	time.Sleep(time.Second)
}

package main

import (
	"fmt"
	"time"
)

func sayHello1() {
	fmt.Println("Hello, World from 1!")
}

func sayHello2() {
	fmt.Println("Hello, World from 2!")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum

}
func main() {
	go sayHello1()
	go sayHello2()
	time.Sleep(1 * time.Second)

	c := make(chan int)
	s := []int{1, 2, 3, 4, 5}
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println("Sum:", x+y)
}

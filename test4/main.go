package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len= %d cap= %d slice=%v\n", len(x), cap(x), x)
}

func sumSlice(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	numbers := make([]int, 0, 5)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers1)

	sum := sumSlice(numbers1)
	fmt.Println("Sum:", sum)
}

package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	fmt.Println("Max value is:", result)
	return result
}

func change(x, y string) (string, string) {
	x, y = y, x
	fmt.Println("change strings:", x, y)
	return x, y
}

func sum(a int) int {
	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}
	fmt.Println("sum is:", sum)
	return sum
}

func main() {
	sum(10)
	max(10, 20)
	change("Hello", "World")
}

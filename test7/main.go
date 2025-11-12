package main

import "fmt"

func change[T any](a, b T) (T, T) {
	return b, a
}

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func remove[T comparable](slice []T, target T) []T {
	result := []T{}
	for _, v := range slice {
		if v != target {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	a, b := 1, 2
	fmt.Println("Before change:", a, b)
	a, b = change(a, b)
	fmt.Println("After change:", a, b)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Contains 3:", Contains(slice, 3))
	fmt.Println("Contains 6:", Contains(slice, 6))
	rmslice := remove(slice, 3)
	fmt.Println("slice after removing :", rmslice)
}

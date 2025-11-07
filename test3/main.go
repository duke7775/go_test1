package main

import "fmt"

type Books struct {
	title  string
	author string
	price  int
}

func arraySum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Array Sum:", sum)
	return sum
}

func change(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func printBooks(book Books) {
	fmt.Println("书名:", book.title)
	fmt.Println("作者:", book.author)
	fmt.Println("价格:", book.price)
}

func main() {
	arraySum([]int{1, 2, 3, 4, 5})

	var a, b int = 10, 20
	change(&a, &b)
	fmt.Println("交换后 a =", a)
	fmt.Println("交换后 b =", b)

	var book1 Books
	book1.title = "沙丘"
	book1.author = "弗兰克"
	book1.price = 100
	printBooks(book1)
}

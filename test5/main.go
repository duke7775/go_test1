package main

import (
	"fmt"
	"strconv"
)

func createMap() map[string]int {
	m := make(map[string]int)
	m["english"] = 90
	m["math"] = 95
	m["science"] = 85
	return m
}
func factory(x int) int {
	if x == 0 {
		return 1
	}
	return x * factory(x-1)
}

func main() {
	scores := createMap()
	fmt.Println(scores)
	delete(scores, "math")
	fmt.Println("删除后的分数:", scores)

	factory(5)
	fmt.Println("5的阶乘是:", factory(5))

	str := "123"
	fmt.Println("字符串内容是:", str)
	num, _ := strconv.Atoi(str)
	fmt.Println("转换后的整数是:", num)

}

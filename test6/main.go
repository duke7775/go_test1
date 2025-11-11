package main

import "fmt"

type cat struct{}

func (c cat) speak() {
	fmt.Println("Cat is speaking")
}

type dog struct{}

func (d dog) speak() {
	fmt.Println("Dog is speaking")
}

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Printf("My name is %s and I am %d years old\n", p.name, p.age)
}

type speaker interface {
	speak()
}

func makeSpeak(s speaker) {
	s.speak()
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "duke", age: 19}
	makeSpeak(p)
	makeSpeak(c)
	makeSpeak(d)
}

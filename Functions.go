package main

import (
	"fmt"
)

func zero() {
	fmt.Println(0)
}

func one(name string) string {
	return "Hello, " + name + "!"
}

func two(firstName, lastName string) string {
	return "My name is " + firstName + " " + lastName
}

func swap(a, b string) (string, string) {
	return b, a
}

func cool(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return
}

func execute(fn func() int) {
	fmt.Println(fn())
}

func main() {
	myFunc := func() int { return 23 }

	execute(myFunc)
}

package main

import (
	"fmt"
)

func basicFor() int {
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}

	return sum
}

func trimmedFor() int {
	num := 1

	for num < 10 {
		num += num
	}

	return num
}

func main() {
	fmt.Println(trimmedFor())
}

package main

import (
	"fmt"
)

func main() {

	var names [2]string

	names[0] = "Derek"
	names[1] = "John"

	namesSlice := names[:]

	namesSlice = append(namesSlice, "Bob", "Jerry", "Mike")

	fmt.Println(namesSlice)
}

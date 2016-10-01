package main

import (
	"fmt"
)

func increment(num *int) {
	*num++
}

func main() {

	num := 50

	pNum := &num

	fmt.Println(num)

}

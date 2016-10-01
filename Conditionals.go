package main

import (
	"fmt"
	"runtime"
)

// basic
func lessThan10(num int) bool {
	if num < 10 {
		return true
	}
	return false
}

// if...else if...else
func greet(time int) string {
	if time < 0 || time > 23 {
		return "Huh?"
	} else if time < 12 {
		return "Good morning"
	} else if time > 12 && time < 17 {
		return "Good afternoon"
	}
	return "Good evening"
}

func platform() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}

func main() {

	platform()
}

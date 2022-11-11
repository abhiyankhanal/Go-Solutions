package main

import "fmt"

func main() {
	input := 3
	result := recursion(input)
	println(result)
}

func recursion(input int) (result int) {
	if input < 0 {
		fmt.Println("The number must be positive")
		return 0
	} else if input <= 1 {
		return 1
	} else {
		return (input * (input - 1))
	}
}

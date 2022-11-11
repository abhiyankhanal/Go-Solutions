package main

import "fmt"

func main() {
	input := 6
	result := recursion(input)
	println(result)
}

func recursion(input int) (result int) {
	if input > 1 {
		return (input * recursion(input-1))
	} else if input >= 0 {
		return 1
	} else {
		fmt.Println("The number must be positive")
		return 0
	}
}

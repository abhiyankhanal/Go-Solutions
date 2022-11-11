package main

import "fmt"

func main() {
	println(isPalindrome(151))
	println(isPalindrome(1010101))

}
func isPalindrome(input int) (result string) {
	var loopVar, reversed int = 0, 0
	original := input
	for input > 0 {
		loopVar = reversed
		rem := input % 10
		reversed = (10 * loopVar) + rem
		input /= 10
	}
	if reversed == original {
		result = fmt.Sprintf("Palindrome, reversed string = %v, input string = %v ", reversed, original)
	} else {
		result = fmt.Sprintf("Not a palindrome, reversed string = %v, input string = %v ", reversed, original)
	}
	return
}

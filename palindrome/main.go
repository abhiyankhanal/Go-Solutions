package main

func main() {
	input := "acaha"
	if isPalindrome(input) == true {
		println("The input string is palindrome")
	} else {
		println("The input string is not a palindrome")
	}

}
func isPalindrome(str string) bool {
	lastIdx := len(str) - 1
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}
	return true
}

package main

func lonelyinteger(a []int32) int32 {
	// Write your code here
	var result int32
	for i := 0; i < len(a); i++ {
		result ^= a[i]
	}
	return result
}

func main() {
	input := []int32{1, 1, 0, 0, 9}
	result := lonelyinteger(input)
	println(result)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{4, 5, 6, 1, 2, 3, 7}
	res := zigzagSequence(input)
	fmt.Println(res)
	//Enter your code here. Read input from STDIN. Print output to STDOUT
}
func zigzagSequence(arr []int) []int {

	sort.Ints(arr)
	maxInt := arr[len(arr)-1]
	midPos := ((len(arr) + 1) / 2) - 1
	// result[midPos] = maxInt
	zig := make([]int, midPos)
	for k := 0; k < midPos; k++ {
		zig[k] = arr[k]
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	zag := make([]int, midPos)
	for m := 0; m < midPos; m++ {
		zag[m] = arr[m+1]
	}
	result := make([]int, 0)
	result = append(result, zig...)
	result = append(result, maxInt)
	result = append(result, zag...)
	return result
}

package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var min, max int64 = 0, 0
	for k := 0; k < 4; k++ {
		max += int64(arr[4-k])
		min += int64(arr[k])
	}
	fmt.Printf("%d %d", min, max)

}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5, 6})
}

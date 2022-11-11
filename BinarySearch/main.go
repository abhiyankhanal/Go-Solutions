package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, l, r, s int) int {

	for l < r {
		m := (l + r) / 2
		if s == arr[m] {
			return m
		} else if arr[m] > s {
			r = m
		} else {
			l = m
		}

	}

	return -1
}

func main() {
	var res []int = []int{3, 4, 5, 6, 7, 8, 1}
	sort.Ints(res)
	l := 0
	r := len(res)
	a := binarySearch(res, l, r, 1)
	if a == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Printf("Found element at pos: %d", a+1)
	}

}

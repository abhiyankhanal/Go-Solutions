package main

import "fmt"

func countingSort(arr []int32) []int32 {
	// Write your code here
	max := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			max = int(arr[i+1])
		}
	}
	var mapArr []int32
	for j := 0; j <= max; j++ {
		mapArr = append(mapArr, 0)
	}
	for k := 0; k < len(arr); k++ {
		mapArr[arr[k]] = mapArr[arr[k]] + 1
	}
	fmt.Printf("%v", mapArr)
	return mapArr
}

func main() {
	input := []int32{1, 1, 1, 2, 3, 4}
	countingSort(input)
}

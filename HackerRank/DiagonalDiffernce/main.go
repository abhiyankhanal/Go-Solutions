package main

import "math"

func diagonalDifference(arr [][]int32) int32 {
	var sumLTR, sumRTL float64 = 0, 0
	// Write your code here
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				sumLTR += float64(arr[i][j])
			}
			if i+j == len(arr)-1 {
				sumRTL += float64(arr[i][j])
			}
		}
	}
	return int32(math.Abs(sumLTR - sumRTL))
}

func main() {
	input := [][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	result := diagonalDifference(input)
	println(result)

}

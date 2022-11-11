package main

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var plus, minus, zero []int32
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			plus = append(plus, arr[i])
		} else if arr[i] < 0 {
			minus = append(minus, arr[i])
		} else {
			zero = append(zero, arr[i])
		}
	}
	size := float32(len(arr))
	var plusRatio, minusRatio, zeroRatio float32 = float32(len(plus)) / size, float32(len(minus)) / size, float32(len(zero)) / size
	fmt.Printf("%.6f\n%.6f\n%.6f", plusRatio, minusRatio, zeroRatio)

}

func main() {
	plusMinus([]int32{3, -1, 2, 0})
}

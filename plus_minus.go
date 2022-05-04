package main

import (
	"fmt"
)

// https://www.hackerrank.com/challenges/plus-minus/problem
// plus_minus([]int32{-4, 3, -9, 0, 4, 1})
func plus_minus(arr []int32) {
	var positiveCounter, negativeCounter, zeroCounter float64
	var size float64 = float64(len(arr))

	for i := range arr {
		if arr[i] > 0 {
			positiveCounter++
		} else if arr[i] < 0 {
			negativeCounter++
		} else {
			zeroCounter++
		}
	}

	fmt.Printf("%.6f\n", positiveCounter/size)
	fmt.Printf("%.6f\n", negativeCounter/size)
	fmt.Printf("%.6f\n", zeroCounter/size)
}

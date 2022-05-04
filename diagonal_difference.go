package main

import (
	"math"
)

// https://www.hackerrank.com/challenges/diagonal-difference/problem
// diagonal_difference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}})
func diagonalDifference(arr [][]int32) int32 {
	var diagonalSum1 int32 = 0
	var diagonalSum2 int32 = 0

	for i, x := range arr {
		diagonalSum1 += x[i]
		diagonalSum2 += x[(len(x)-i)-1]
	}

	return int32(math.Abs(float64(diagonalSum2 - diagonalSum1)))
}

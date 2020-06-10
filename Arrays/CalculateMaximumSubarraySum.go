package main

import (
	"fmt"
	"math"
)

// Given an array of numbers, find the maximum sum of any contiguous subarray of
// the array. For example, given the array [34, -50, 42, 14, -5, 86], the maximum sum
// would be 137, since we would take elements 42, 14, -5,  and 86.
// Given the array [-5, -1, -8, -9], the maximum sum would be 0, since we would choose not to
// take any elements.

// using kadane's algorithm


func main() {

	fmt.Print(maximumSum([]int{34, -50, 42, 14, -5, 86}))
}

func maximumSum(arr []int) int {

	var max, maxSum float64

	for _, elem := range arr {
		maxSum = math.Max(float64(elem), maxSum + float64(elem))
		max = math.Max(max, maxSum)
	}

	return int(max)
}

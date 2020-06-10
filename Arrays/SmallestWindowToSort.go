package main

import (
	"fmt"
	"math"
)

// Given an array of integers that are out of order,
// determine the bounds of the smallest window that must be
// sorted in order for the entire array to be sorted.
// For example, given [3, 7, 5, 6, 9], you should return (1, 3)

func main() {

	fmt.Print(window([]int{3, 6, 10, 9, 11, 4, 12}))
}

func window(arr []int) (int, int) {
	var right, left int

	maxInt64, minInt64 := math.MaxInt64, math.MinInt64


	for index := len(arr) - 1; index >= 0; index-- {

		if maxInt64 < arr[index] {
			left = index
		}
		maxInt64 = arr[index]
	}

	for index, elem := range arr {

		if minInt64 > elem {
			right = index
		} else {
			minInt64 = elem
		}

	}

	return left, right
}
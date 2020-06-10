package main

import "fmt"

// Given an array of integers, return a new array such that each element at
// index i of the new array is the product of all the numbers in the original array except the one at i.

// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
// If our input was [3, 2, 1], the expected output would be [2, 3, 6].

func main() {

	fmt.Print(products([]int{1, 2, 3, 4, 5}))
}

func products(nums []int) []int {

	lenOfNums := len(nums)

	// prefix numbers
	prefixNumbers := make([]int, lenOfNums)
	for index, num := range nums {

		if index != 0 {
			prefixNumbers[index] = prefixNumbers[index - 1] * num
		} else {
			prefixNumbers[index] = num
		}
	}

	// suffix numbers
	suffixNumber := make([]int, lenOfNums)
	last := lenOfNums - 1
	for index := range nums {

		if index != 0 {
			suffixNumber[last - index] = suffixNumber[last - index + 1] * nums[last - index]
		} else {
			suffixNumber[last] = nums[last]
		}

	}

	// result calculation
	result := make([]int, lenOfNums)
	for index := range nums {
		if index == 0 {
			result[index] = suffixNumber[index + 1]
		} else if index == last {
			result[index] = prefixNumbers[index - 1]
		} else {
			result[index] = prefixNumbers[index - 1] * suffixNumber[index + 1]
 		}
	}
	return result
}
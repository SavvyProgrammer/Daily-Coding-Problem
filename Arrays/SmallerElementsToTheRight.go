package main

import (
	"fmt"
	"sort"
)

/*

Given an array of integers, return a new array where each element in the new array
is the number of smaller elements to the right of that element in the original input
array.

For example, given the array [3, 4, 9, 6, 1], return [1. 1. 2, 1, 0], since:

There is 1 smaller element to the right of 3
There is 1 smaller element to the right of 4
There are 2 smaller elements to the right of 9
There is 1 smaller element to the right of 6
There are no smaller elements to the right of 1

*/


func main() {

	fmt.Print(smallerElements([]int{3, 4, 9, 6, 1}))
}


func smallerElements(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	var arrToSort []int
	lastIndex := n - 1
	for index := range arr {
		currentIndex := lastIndex - index
		elem := arr[currentIndex]
		result[currentIndex] = sort.SearchInts(arrToSort, elem)
		arrToSort = append(arrToSort, elem)
	}
	return result
}
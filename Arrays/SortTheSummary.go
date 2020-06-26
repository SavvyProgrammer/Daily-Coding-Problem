package main
// Sort the Summary Hackerrank Golang Solution
// Given an array of integers, create a 2-dimensional array where the first element is a distinct value from
// the array and the second element is that value's frequency within the array.
// Sort the resulting array descending by frequency. If multiple values have the same frequency; they should
// be sorted ascending.

import (
	"fmt"
)

func main() {
	fmt.Println(groupSort([]int32{3, 3, 1, 2, 1, 4, 4, 5, 6, 9, 9, 7, 5}))
}
func groupSort(arr []int32) [][]int32 {
	counter := make(map[int32]int32)
	for _, elem := range arr {
		counter[elem]++
	}
	result := make([][]int32, len(counter))
	var indexOfResult int
	for key, value := range counter {
		sliceOfKeyValue := make([]int32, 2)
		sliceOfKeyValue[0], sliceOfKeyValue[1] = key, value
		result[indexOfResult] = sliceOfKeyValue
		indexOfResult++
	}
	for i := 1; i < len(result); i++ {
		key := result[i][0]
		freq := result[i][1]
		slice := result[i]
		j := i - 1
		for j > -1 && freq > result[j][1] {
			result[j+1] = result[j]
			j--
		}
		for j >= 0 && result[j][0] > key && freq == result[j][1] {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = slice
	}
	return result
}

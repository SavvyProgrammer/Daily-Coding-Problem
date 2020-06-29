package main

import "fmt"

// Given a word w and a string s, find all indices in s which are the starting location of anagrams of w.
// For example, given w is ab and s is abxaba, return [0, 3, 4].
// For an optimal solution, I'll use Rabin Karp Algorithm
func main() {
	fmt.Println(anagramIndices("abcd", "bacdgabcda"))
}

func anagramIndices(word, s string) []int {
	var result []int
	freq := make(map[uint8]int)

	for index := range word {
		freq[s[index]]++
	}

	for index := range s[:len(word)] {
		freq[s[index]]--
		deleteIfZero(freq, s[index])
	}

	if len(freq) == 0 {
		result = append(result, 0)
	}

	for index := len(word); index < len(s); index++ {
		firstChar, endChar := s[index-len(word)], s[index]
		freq[firstChar]++
		deleteIfZero(freq, firstChar)
		freq[endChar]--
		deleteIfZero(freq, endChar)
		if len(freq) == 0 {
			result = append(result, index-len(word)+1)
		}
	}

	return result
}

func deleteIfZero(freq map[uint8]int, elem uint8) {
	if freq[elem] == 0 {
		delete(freq, elem)
	}
}

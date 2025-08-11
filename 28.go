// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// Example 1:

// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:

// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

package main

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)

	for i := 0; i <= n-m; i++ {
		countie := 0
		for j := 0; j < m; j++ {
			if haystack[i+j] == needle[j] {
				countie++
			} else {
				break
			}
		}
		if countie == m {
			return i
		}
	}
	return -1
}


func main() {
	h1 := "hello"
	n1 := "ll"

	fmt.Println(strStr(h1, n1))
}
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return "" 
	}
	prefix := strs[0]

	for _, s := range strs[1:] { 
		for len(prefix) > 0 && len(s) < len(prefix) || s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1] 
		}
		if prefix == "" {
			return "" 
		}
	}

	return prefix
}

func main() {
	s1 := []string { "flower","flow","flight"}
	s2 := []string {"dog","racecar","car"}

	fmt.Println(longestCommonPrefix(s1))
	fmt.Println(longestCommonPrefix(s2))
}


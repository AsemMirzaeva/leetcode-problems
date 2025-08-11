// 2185. Counting Words With a Given Prefix

// You are given an array of strings words and a string pref.

// Return the number of strings in words that contain pref as a prefix.

// A prefix of a string s is any leading contiguous substring of s.

// Example 1:

// Input: words = ["pay","attention","practice","attend"], pref = "at"
// Output: 2
// Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
// Example 2:

// Input: words = ["leetcode","win","loops","success"], pref = "code"
// Output: 0
// Explanation: There are no strings that contain "code" as a prefix.

package main


import "fmt"

func prefixCount(words []string, pref string) int {

	count := 0

	for idx, _ := range words {

		length := len(pref)

		if length <= len(words[idx]) && words[idx][:length] == pref {
			count++
		}
	}
	return  count
    
}


func main() {
	a1 := []string {"pay","attention","practice","attend"}
	p1 := "at"
	a2 := []string {"leetcode","win","loops","success"}
	p2 := "code"

	fmt.Println(prefixCount(a1,p1))
	fmt.Println(prefixCount(a2,p2))
}
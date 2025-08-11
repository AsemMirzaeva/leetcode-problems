// Given a string s, return the number of unique palindromes of length three that are a subsequence of s.

// Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.

// A palindrome is a string that reads the same forwards and backwards.

// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.

// For example, "ace" is a subsequence of "abcde".

// Example 1:

// Input: s = "aabca"
// Output: 3
// Explanation: The 3 palindromic subsequences of length 3 are:
// - "aba" (subsequence of "aabca")
// - "aaa" (subsequence of "aabca")
// - "aca" (subsequence of "aabca")
// Example 2:

// Input: s = "adc"
// Output: 0
// Explanation: There are no palindromic subsequences of length 3 in "adc".
// Example 3:

// Input: s = "bbcbaba"
// Output: 4
// Explanation: The 4 palindromic subsequences of length 3 are:
// - "bbb" (subsequence of "bbcbaba")
// - "bcb" (subsequence of "bbcbaba")
// - "bab" (subsequence of "bbcbaba")
// - "aba" (subsequence of "bbcbaba")

package main

import (
	"fmt"
)

// func countPalindromicSubsequence(s string) int {

// 	// unique := make(map[string]bool)
// 	// var result []string


// 	// for i := 1; i < len(s)-1; i++ {
// 	// 	for left := 0; left < i; left++ {
// 	// 		for right := i + 1; right < len(s); right++ {
// 	// 			if s[left] == s[right] {
					
// 	// 				triplet := string(s[left]) + string(s[i]) + string(s[right])
// 	// 				if !unique[triplet] {
// 	// 					unique[triplet] = true
// 	// 					result = append(result, triplet)
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	// return len(result)
// }



// // optimized

// func countPalindromicSubsequence(s string) int {
//     indexes := make(map[rune][]int)
//     for k, v := range s {
//         if indexes[v] == nil {
//             indexes[v] = make([]int, 0)
//         }
//         indexes[v] = append(indexes[v], k)
//     }

//     res := 0
//     for _, idxs := range indexes {
//         if len(idxs) >= 2 {
//             l := idxs[0]
//             r := idxs[len(idxs)-1]

//             for _, idxs2 := range indexes {
//                 for _, i := range idxs2 {
//                     if i > l && i < r  {
//                         res++
//                         break
//                     }
//                 }
//             }
//         }
//     }
//     return res
// }


func countPalindromicSubsequence(s string) int {
    left := make(map[byte]bool)  
    result := make(map[string]bool) 

    rightCount := make([]int, 26) 
    for i := 0; i < len(s); i++ {
        rightCount[s[i]-'a']++
    }

    for i := 0; i < len(s); i++ {
        rightCount[s[i]-'a']-- 
        for ch := byte('a'); ch <= 'z'; ch++ {
            if left[ch] && rightCount[ch-'a'] > 0 {
                triplet := string(ch) + string(s[i]) + string(ch)
                result[triplet] = true
            }
        }
        left[s[i]] = true 
    }

    return len(result)
}

func main() {
	s := "bbcbaba"

	fmt.Println(countPalindromicSubsequence(s))
}

// Given a string s and an integer k, return true if you can use all the characters
//  in s to construct k palindrome strings or false otherwise.

// Example 1:

// Input: s = "annabelle", k = 2
// Output: true
// Explanation: You can construct two palindromes using all characters in s.
// Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"
// Example 2:

// Input: s = "leetcode", k = 3
// Output: false
// Explanation: It is impossible to construct 3 palindromes using all the characters of s.
// Example 3:

// Input: s = "true", k = 4
// Output: true
// Explanation: The only possible solution is to put each character in a separate string.

package main

import "fmt"

// func countPalindromicSubsequence(s string) int {
//     left := make(map[byte]bool)
//     result := make(map[string]bool)

//     rightCount := make([]int, 26)
//     for i := 0; i < len(s); i++ {
//         rightCount[s[i]-'a']++
//     }

//     for i := 0; i < len(s); i++ {
//         rightCount[s[i]-'a']--
//         for ch := byte('a'); ch <= 'z'; ch++ {
//             if left[ch] && rightCount[ch-'a'] > 0 {
//                 triplet := string(ch) + string(s[i]) + string(ch)
//                 result[triplet] = true
//             }
//         }
//         left[s[i]] = true
//     }

//     return len(result)
// }


func canConstruct(s string, k int) bool {
	if k > len(s) {
		return false
	}

	charCount := make(map[rune]int)
	for _, ch := range s {
		charCount[ch]++
	}

	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}


	return oddCount <= k
}


func canConstruct(s string, k int) bool {
    if len(s) < k {
        return false
    }
    m := make([]int, 26)
    for _, c := range s {
        m[c - 'a']++
    } 
    n := 0
    for _, v := range m {
        if v % 2 == 1 {
            n++
        }
    }
    if n > k {
        return false
    }
    return true
}

func main() {
	s1 := "annabelle"
	k1 := 2

	s2 := "leetcode"
	k2 := 3

	s3 := "true"
	k3 := 4

	fmt.Println(canConstruct(s1, k1))
	fmt.Println(canConstruct(s2, k2))
	fmt.Println(canConstruct(s3, k3))
}

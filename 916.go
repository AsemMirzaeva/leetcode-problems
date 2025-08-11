// You are given two string arrays words1 and words2.

// A string b is a subset of string a if every letter in b occurs in a including multiplicity.

// For example, "wrr" is a subset of "warrior" but is not a subset of "world".
// A string a from words1 is universal if for every string b in words2, b is a subset of a.

// Return an array of all the universal strings in words1. You may return the answer in any order.

// Example 1:

// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
// Output: ["facebook","google","leetcode"]
// Example 2:

// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
// Output: ["apple","google","leetcode"]

package main

import "fmt"


// // optimized 

// type cFreq [26]byte

// func wordSubsets(candidates, targets []string) (subs []string) {
// 	target := reduce(targets)
// 	for i := range candidates {
// 		if !isSubset(&candidates[i], &target) {
// 			continue
// 		}
// 		subs = append(subs, candidates[i])
// 	}

// 	return
// }

// func reduce(targets []string) (word string) {
// 	cmap := cFreq{}

// 	for i := range targets {
// 		tmap := cFreq{}
// 		for _, c := range targets[i] {
// 			tmap[c-'a']++
// 		}
// 		for i, count := range tmap {
// 			cmap[i] = max(cmap[i], count)
// 		}
// 	}

// 	for i, count := range cmap {
// 		for range count {
// 			word += string(byte(i + 'a'))
// 		}
// 	}

// 	return word
// }

// func isSubset(candidate, target *string) bool {
// 	cmap := cFreq{}

// 	for _, c := range *candidate {
// 		cmap[c-'a']++
// 	}

// 	for _, c := range *target {
// 		if cmap[c-'a'] == 0 {
// 			return false
// 		}

// 		cmap[c-'a']--
// 	}

// 	return true
// }

func wordSubsets(words1 []string, words2 []string) []string {
	word2Freq := make(map[rune]int)
	for _, word := range words2 {

		tempFreq := make(map[rune]int)
		for _, ch := range word {
			tempFreq[ch]++
		}
	
		for ch, count := range tempFreq {
			if count > word2Freq[ch] {
				word2Freq[ch] = count
			}
		}
	}

	var result []string
	for _, word := range words1 {
		wordFreq := make(map[rune]int)

		for _, ch := range word {
			wordFreq[ch]++
		}

		isValid := true
		for ch, count := range word2Freq {
			if wordFreq[ch] < count {
				isValid = false
				break
			}
		}
		if isValid {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	w1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	s1 := []string{"e", "o"}

	w2 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	s2 := []string{"l", "e"}

	fmt.Println(wordSubsets(w1, s1))
	fmt.Println(wordSubsets(w2, s2))
}

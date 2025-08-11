// You are given a 0-indexed array of strings words and a 2D array of integers queries.

// Each query queries[i] = [li, ri] asks us to find the number of strings present in the range li to ri (both inclusive) of words that start and end with a vowel.

// Return an array ans of size queries.length, where ans[i] is the answer to the ith query.

// Note that the vowel letters are 'a', 'e', 'i', 'o', and 'u'.

 

// Example 1:

// Input: words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
// Output: [2,3,0]
// Explanation: The strings starting and ending with a vowel are "aba", "ece", "aa" and "e".
// The answer to the query [0,2] is 2 (strings "aba" and "ece").
// to query [1,4] is 3 (strings "ece", "aa", "e").
// to query [1,1] is 0.
// We return [2,3,0].
// Example 2:

// Input: words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
// Output: [3,2,1]
// Explanation: Every string satisfies the conditions, so we return [3,2,1].


package main





// func vowelStrings(words []string, queries [][]int) []int {
// 	vowels := map[rune]bool{'a':true, 'e':true, 'i': true, 'o':true, 'u':true}

// 	isVowel := func(ch rune) bool {
// 		_, exists := vowels[ch]
// 		return exists
// 	}
// 	results := []int{}


// 	for _, query := range queries {
// 		start, end := query[0], query[1]
// 		count := 0

// 		for i := start; i <= end; i++ {
// 			value := words[i]
// 			first := rune(value[0])
// 			last := rune(value[len(value)-1])

// 			switch {
// 			case isVowel(first) && isVowel(last):
// 				count++
// 			}
// 		}
// 		results = append(results, count)
// 	}
// 	return results
	 
// }



// func vowelStrings(words []string, queries [][]int) []int {
// 	vowels := map[rune]bool{'a':true, 'e':true, 'i': true, 'o':true, 'u':true}

// 	isVowel := func(ch rune) bool {
// 		_, exists := vowels[ch]
// 		return exists
// 	}

	
// 	n := len(words)
// 	prefixSum := make([]int, n+1)

// 	for i, word := range words {
// 		first := rune(word[0])
// 		last := rune(word[len(word)-1])

// 		if isVowel(first) && isVowel(last) {
// 			prefixSum[i+1] = prefixSum[i] + 1
// 		} else {
// 			prefixSum[i+1] = prefixSum[i]
// 		}
// 	}


// 	results := []int{}


// 	for _, query := range queries {
// 		start, end := query[0], query[1]
// 		count := prefixSum[end+1] - prefixSum[start]
// 		results = append(results, count)

		
// 	}
// 	return results
	 
// }



func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func vowelStrings(words []string, queries [][]int) []int {
    n := len(words)
    prefix := make([]int, n)

    for idx, word := range(words) {
        prefix[idx] = 0
        if idx > 0 {
            prefix[idx] = prefix[idx-1]
        }

        if isVowel(word[0]) && isVowel(word[len(word) - 1]) {
            prefix[idx]++
        }
    }    
    sols := make([]int, len(queries))
    for idx, query := range(queries) {
        l := query[0]
        r := query[1]
        if l == 0 {
            sols[idx] = prefix[r]
            continue
        }
        sols[idx] = prefix[r] - prefix[l - 1]
    }
    return sols
}

func main() {



}
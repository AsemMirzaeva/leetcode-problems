// You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

// Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.

// Example 1:

// Input: s1 = "bank", s2 = "kanb"
// Output: true
// Explanation: For example, swap the first character with the last character of s2 to make "bank".
// Example 2:

// Input: s1 = "attack", s2 = "defend"
// Output: false
// Explanation: It is impossible to make them equal with one string swap.
// Example 3:

// Input: s1 = "kelb", s2 = "kelb"
// Output: true
// Explanation: The two strings are already equal, so no string swap operation is required.

package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	// n := len(s1)
	// if s1 == s2 {
	// 	return true
	// }
	
    // for i := 0; i < n; i++ {
	// 	runes := []rune(s1)

	// 	runes[i], runes[n-i-1] = runes[n-i-1], runes[i] 

	// 	s11 := string(runes)

	// 	if s11 == s2 {
	// 		return true
	// 	} 

	// 	for j := i + 1; j < n; j++ {

    //         runes[i], runes[j] = runes[j], runes[i]
    //         s11 := string(runes)

    //         if s11 == s2 {
    //             return true
    //         }

        
    //         runes[i], runes[j] = runes[j], runes[i]
    //     }
	// }
	// return false

	n := len(s1)
    if s1 == s2 {
        return true
    }

    runes := []rune(s1) 

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
    
            runes[i], runes[j] = runes[j], runes[i]
            s11 := string(runes)

            if s11 == s2 {
                return true
            }

            runes[i], runes[j] = runes[j], runes[i]
        }
    }
    
    return false
}

func main() {
	s1 := "bank"
	s2 := "kanb"

	s3 := "attack"
	s4 := "defend"

	s5 := "abcd"
	s6 := "dcba"

	fmt.Println(areAlmostEqual(s1, s2))
	fmt.Println(areAlmostEqual(s3, s4))
	fmt.Println(areAlmostEqual(s5, s6))


}
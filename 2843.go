// You are given two positive integers low and high.

// An integer x consisting of 2 * n digits is symmetric if the sum of the first n digits of x is equal to the sum of the last n digits of x. Numbers with an odd number of digits are never symmetric.

// Return the number of symmetric integers in the range [low, high].

// Example 1:

// Input: low = 1, high = 100
// Output: 9
// Explanation: There are 9 symmetric integers between 1 and 100: 11, 22, 33, 44, 55, 66, 77, 88, and 99.
// Example 2:

// Input: low = 1200, high = 1230
// Output: 4
// Explanation: There are 4 symmetric integers between 1200 and 1230: 1203, 1212, 1221, and 1230.

// 10 000 10-100 1000-9999
package main

import (
	"fmt"
)

func countSymmetricIntegers(low int, high int) int {
	countie := 0

	for i := low; i <= high; i++ {
		if i <= 100 {
			firstH := i / 10
			secondH := i % 10
			if firstH == secondH {
				countie++
	
			}
		} else if i >= 1000 {
			firstH := getSumDigit(i / 100)
			secondH := getSumDigit(i % 100)
			if firstH == secondH {
				countie++
			}
		}
	}

	return countie
}

func getSumDigit(n int) int {
	return n/10 + n%10
}

func main() {
	l1 := 1
	h1 := 100

	l2 := 1200
	h2 := 1230

	l3 := 100
	h3 := 1782

	fmt.Println(countSymmetricIntegers(l1, h1))
	fmt.Println(countSymmetricIntegers(l2, h2))
	fmt.Println(countSymmetricIntegers(l3, h3))
}

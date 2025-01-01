// Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).

// The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.

// Example 1:

// Input: s = "011101"
// Output: 5
// Explanation:
// All possible ways of splitting s into two non-empty substrings are:
// left = "0" and right = "11101", score = 1 + 4 = 5
// left = "01" and right = "1101", score = 1 + 3 = 4
// left = "011" and right = "101", score = 1 + 2 = 3
// left = "0111" and right = "01", score = 1 + 1 = 2
// left = "01110" and right = "1", score = 2 + 1 = 3
// Example 2:

// Input: s = "00111"
// Output: 5
// Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
// Example 3:

// Input: s = "1111"
// Output: 3


package main

import (
	"fmt"
	"strconv"
)

func maxScore(s string) int {

	var nums, sum []int
	

	for _, ch := range s {
		str := string(ch)
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	for n := 0; n < len(nums); n++{
		count0, count1 := 0, 0

	    counter := n + 1

		s1 := []int{}
		for k := 0; k < counter && k < len(nums); k++ {
			s1 = append(s1, nums[k])
			if nums[k] == 0 {
				count0 ++
			}
		}
		
		s2 := []int{}
		for m := counter; m < len(nums); m++ {
			s2 = append(s2, nums[m])
			if nums[m] == 1 {
				count1 ++
			}
		}
		if len(s2) > 0 {
			sum = append(sum, count0 + count1)
		}
		
	}


	if len(sum) == 0 {
		return 0
	}

	max := sum[0]
	for _, num := range sum {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	s := "00"

	fmt.Println(maxScore(s))
}

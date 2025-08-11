// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:

// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109

package main

import "fmt"

// Using Boyer-Moore Voting Algorithm for better efficiency
func majorityElement(nums []int) int {
	// Candidate selection phase
	candidate := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	nums := []int{3, 2, 3, 3, 1, 3}
	fmt.Println("Majority Element:", majorityElement(nums)) // Output: 3
}

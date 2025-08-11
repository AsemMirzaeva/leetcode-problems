// You have a set of integers s, which originally contains
// all the numbers from 1 to n. Unfortunately, due to some error,
// one of the numbers in s got duplicated to another number in the set,
//  which results in repetition of one number and loss of another number.

// You are given an integer array nums representing the data status of this set
//  after the error.

// Find the number that occurs twice and the number
// that is missing and return them in the form of an array.

// Example 1:

// Input: nums = [1,2,2,4]
// Output: [2,3]
// Example 2:

// Input: nums = [1,1]
// Output: [1,2]

package main

import "fmt"

func findErrorNums(nums []int) []int {
	seen := make(map[int]bool)
	var duplicate, missing int

	for i := 1; i <= len(nums); i++ {
		if _, exists := seen[nums[i-1]]; exists {
			duplicate = nums[i-1]
		} else {
			seen[nums[i-1]] = true
		}
	}

	for i := 1; i <= len(nums); i++ {
		if !seen[i] {
			missing = i
			break
		}
	}
	return []int{duplicate, missing}
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}

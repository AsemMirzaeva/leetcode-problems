// Given an array nums, return true if the array was originally sorted in non-decreasing order, then rotated some number of positions (including zero). Otherwise, return false.

// There may be duplicates in the original array.

// Note: An array A rotated by x positions results in an array B of the same length such that A[i] == B[(i+x) % A.length], where % is the modulo operation.

// Example 1:

// Input: nums = [3,4,5,1,2]
// Output: true
// Explanation: [1,2,3,4,5] is the original sorted array.
// You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2].
// Example 2:

// Input: nums = [2,1,3,4]
// Output: false
// Explanation: There is no sorted array once rotated that can make nums.
// Example 3:

// Input: nums = [1,2,3]
// Output: true
// Explanation: [1,2,3] is the original sorted array.
// You can rotate the array by x = 0 positions (i.e. no rotation) to make nums.

package main

import "fmt"

func check(nums []int) bool {
    count := 0 
	
	n := len(nums)

	
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			count ++
		}
		if count > 1 {
			return false
		}
	
	}

	return true

}


func main() {
	n1 := []int {3,4,5,1,2}
	n2 := []int {2,1,3,4}
	n3 := []int {1,2,3}

	fmt.Println(check(n1))
	fmt.Println(check(n2))
	fmt.Println(check(n3))
}
// Given an integer array nums, return the number of subarrays of length 3 such that the sum of the first and third numbers equals exactly half of the second number.

// Example 1:

// Input: nums = [1,2,1,4,1]

// Output: 1

// Explanation:

// Only the subarray [1,4,1] contains exactly 3 elements where the sum of the first and third numbers equals half the middle number.

// Example 2:

// Input: nums = [1,1,1]

// Output: 0

// Explanation:

// [1,1,1] is the only subarray of length 3. However, its first and third numbers do not add to half the middle number.

package main

import "fmt"

func countSubarrays(nums []int) int {
	n := len(nums)

	if n < 3 {
		return 0
	}
	if n == 3 {
		if (nums[0] + nums[2]) == nums[1]/2 {
			return 1
		}
		return 0
	}

	var subAr [][] int
	
	for i :=0; i < n -2; i ++ {
		if float32(nums[i] + nums[i+2]) == float32(nums[i+1])/2 {
			fmt.Println(nums[i+1]/2)
			subAr = append(subAr, []int{nums[i], nums[i+1], nums[i+2]})
		}
	}
	fmt.Println(subAr)
	return len(subAr)
}

func main() {
	n1 := []int {-1,-4,-1,4}
	// n2 := []int {1,1,1}

	fmt.Println(countSubarrays(n1))
	// fmt.Println(countSubarrays(n2))
}
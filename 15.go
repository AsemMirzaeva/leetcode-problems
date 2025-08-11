// 15. 3Sum
// Medium

// Topics
// Companies

// Hint
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Example 2:

// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// Example 3:

// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

package main

import (
	"fmt"
	"sort"
)



func threeSum(nums []int) [][]int {
	
	// var triplets [][]int
	// seen := make(map[string]bool)


	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums)-1; j++ {
	// 		for k := j + 1; k < len(nums); k++ {
	// 			if nums[i] + nums[j] + nums[k] == 0 {
	// 				key := fmt.Sprintf("%d,%d,%d", nums[i], nums[j], nums[k])
	// 				if !seen[key] {
	// 					triplets = append(triplets, []int{nums[i], nums[j], nums[k]})
	// 					seen[key] = true
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// return triplets

	var triplets [][]int
	// First, sort the numbers
	sort.Ints(nums)

	// Iterate through the numbers
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate values of `nums[i]` to avoid duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two pointers to find pairs that sum to -nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// We found a valid triplet
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate values for `left` and `right`
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move the pointers
				left++
				right--
			} else if sum < 0 {
				left++ // We need a larger number, so move the `left` pointer
			} else {
				right-- // We need a smaller number, so move the `right` pointer
			}
		}
	}

	// Return the resulting triplets
	return triplets
}

func main() {
	n1 := []int {-1,0,1,2,-1,-4}
	n2 := []int {0,1,1}
	n3 := []int {0,0,0}
	n4 := []int{0,0,0,0}

	fmt.Println(threeSum(n1))
	fmt.Println(threeSum(n2))
	fmt.Println(threeSum(n3))
	fmt.Println(threeSum(n4))
}
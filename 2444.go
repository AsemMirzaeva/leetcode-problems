// You are given an integer array nums and two integers minK and maxK.

// A fixed-bound subarray of nums is a subarray that satisfies the following conditions:

// The minimum value in the subarray is equal to minK.
// The maximum value in the subarray is equal to maxK.
// Return the number of fixed-bound subarrays.

// A subarray is a contiguous part of an array.

// Example 1:

// Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
// Output: 2
// Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
// Example 2:

// Input: nums = [1,1,1,1], minK = 1, maxK = 1
// Output: 10
// Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.

package main

import "fmt"

func countSubarrays(nums []int, minK int, maxK int) int64 {

	// 1, 3, 5, 2, 7, 5

	var ans int64
	lastMinK, lastMaxK := -1, -1
	lastInvalid := -1

	for i, v := range nums {
		if v < minK || v > maxK {
			lastInvalid = i
		}
		if v == minK {
			lastMinK = i
		}
		if v == maxK {
			lastMaxK = i
		}
	
		validStart := min(lastMinK, lastMaxK)
		if validStart > lastInvalid {
			ans += int64(validStart - lastInvalid)
		}
	}

	return ans
}


func countSubarrays(nums []int, minK int, maxK int) int64 {
	// nums = append(nums, 0)
	s := 1<<31 - 1
	minFound := 1<<31 - 1
	maxFound := 1<<31 - 1
    ans := 0
	for i := range len(nums) {
		if nums[i] < minK || nums[i] > maxK {
			minFound = 1<<31 - 1
			maxFound = 1<<31 - 1
            s = 1<<31 - 1
		} else {
			if nums[i] == maxK {
				maxFound = i
			}
			if nums[i] == minK {
				minFound = i
			}
            s = min(s, i)
            if maxFound < 1<<31 - 1 && minFound < 1<<31 - 1 {
                ans += min(maxFound, minFound) - s + 1
            }
		}
	}
    return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	m1 := 1
	m2 := 5
	n1 := []int{1, 3, 5, 2, 7, 5}

	m3 := 1
	m4 := 1
	n2 := []int{1, 1, 1, 1}

	fmt.Println(countSubarrays(n1, m1, m2))
	fmt.Println(countSubarrays(n2, m3, m4))

}

// You are given an integer array nums and a positive integer k.

// Return the number of subarrays where the maximum element of nums appears at least k times in that subarray.

// A subarray is a contiguous sequence of elements within an array.

// Example 1:

// Input: nums = [1,3,2,3,3], k = 2
// Output: 6
// Explanation: The subarrays that contain the element 3 at least 2 times are: [1,3,2,3], [1,3,2,3,3], [3,2,3], [3,2,3,3], [2,3,3] and [3,3].
// Example 2:

// Input: nums = [1,4,2,1], k = 3
// Output: 0
// Explanation: No subarray contains the element 4 at least 3 times.

package main

import "fmt"

func countSubarrays(nums []int, k int) int64 {
    maxVal := max(nums)
    var  result int64 = 0
    left := 0
    maxCount := 0

    for right := 0; right < len(nums); right++ {
        if nums[right] == maxVal {
            maxCount++
        }

        // Move left forward until we have at least k maxVal's in the window
        for maxCount >= k {
            if nums[left] == maxVal {
                maxCount--
            }
            left++
        }

        // All subarrays ending at `right` and starting from 0 to `left-1`
        result += int64(left)
    }

    return result
}

func max(nums []int) int {
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
	n1 := []int{1, 3, 2, 3, 3}
	k1 := 2

	// n2 := []int{1, 4, 2, 1}
	// k2 := 3

	fmt.Println(countSubarrays(n1, k1))
	// fmt.Println(countSubarrays(n2, k2))
}

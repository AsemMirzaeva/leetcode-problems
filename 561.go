// Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) 
// such that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.

 

// Example 1:

// Input: nums = [1,4,3,2]
// Output: 4
// Explanation: All possible pairings (ignoring the ordering of elements) are:
// 1. (1, 4), (2, 3) -> min(1, 4) + min(2, 3) = 1 + 2 = 3
// 2. (1, 3), (2, 4) -> min(1, 3) + min(2, 4) = 1 + 2 = 3
// 3. (1, 2), (3, 4) -> min(1, 2) + min(3, 4) = 1 + 3 = 4
// So the maximum possible sum is 4.
// Example 2:

// Input: nums = [6,2,6,5,1,2]
// Output: 9
// Explanation: The optimal pairing is (2, 1), (2, 5), (6, 6). min(2, 1) + min(2, 5) + min(6, 6) = 1 + 2 + 6 = 9.


func arrayPairSum(nums []int) int {
    n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	pairs := make ([][2]int, 0)
	for i := 0; i < n-1; i +=2 {
		pairs = append(pairs, [2]int{nums[i], nums[i+1]})
	}

	sumOfKeys := 0
	for _, pair := range pairs {
		sumOfKeys += pair[0]
	}

	return sumOfKeys
}

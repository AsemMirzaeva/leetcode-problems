// Given an array of integers nums, half of the integers in nums are odd, and the other half are even.

// Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.

// Return any answer array that satisfies this condition.

 

// Example 1:

// Input: nums = [4,2,5,7]
// Output: [4,5,2,7]
// Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
// Example 2:

// Input: nums = [2,3]
// Output: [2,3]


func sortArrayByParity(nums []int) []int {
	n := len(nums)
	paired := make([]int, n) 

	evenIndex, oddIndex := 0, 1 

	for i := 0; i < n; i++ {
		if nums[i] % 2 == 0 {
			paired[evenIndex] = nums[i]
			evenIndex += 2
		} else {
			paired[oddIndex] = nums[i]
			oddIndex += 2
		}
	}
	return paired
}
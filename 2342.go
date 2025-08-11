// You are given a 0-indexed array nums consisting of positive integers. You can choose two indices i and j, such that i != j, and the sum of digits of the number nums[i] is equal to that of nums[j].

// Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and j that satisfy the conditions.

// Example 1:

// Input: nums = [18,43,36,13,7]
// Output: 54
// Explanation: The pairs (i, j) that satisfy the conditions are:
// - (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
// - (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
// So the maximum sum that we can obtain is 54.
// Example 2:

// Input: nums = [10,12,19,14]
// Output: -1
// Explanation: There are no two numbers that satisfy the conditions, so we return -1.

package main

import "fmt"

func sumOfDigits(n int) int {
	sum := 0 
	for n > 0 {
		sum += n%10
		n /= 10
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumSum(nums []int) int {
    digitSumMap := make(map[int][]int)
	
	for _, num := range nums {
		digitSum := sumOfDigits(num)
		digitSumMap[digitSum] = append(digitSumMap[digitSum], num)
	}

	maxSum := -1

	for _, values := range digitSumMap {
		if len(values) > 1 {
			firstmax, secondMax := 0,0
			for _, v := range values {
				if v > firstmax {
					secondMax = firstmax
					firstmax = v
				} else if v > secondMax {
					secondMax = v
				}
			}
			maxSum = max(maxSum, firstmax+secondMax)
		}
	}
	return maxSum 
}

func main() {
	n1 := []int {18,43,36,13,7}

	n2 := []int {10,12,19,14}

	fmt.Println(maximumSum(n1))
	fmt.Println(maximumSum(n2))
}
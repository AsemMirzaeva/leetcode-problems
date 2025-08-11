// Given an array nums of distinct positive integers,
// return the number of tuples (a, b, c, d) such
// that a * b = c * d where a, b, c, and d are elements of nums, and a != b != c != d.

// Example 1:

// Input: nums = [2,3,4,6]
// Output: 8
// Explanation: There are 8 valid tuples:
// (2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
// (3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
// Example 2:

// Input: nums = [1,2,4,5,10]
// Output: 16
// Explanation: There are 16 valid tuples:
// (1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
// (2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
// (2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,5,4)
// (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)

package main

import "fmt"

func tupleSameProduct(nums []int) int {
    var tuples [][]int
	n := len(nums) 
	for i := 0; i < n-1; i++ { 
		for j := 0; j < n; j++ { 
			if i == j { continue } 
			
			for k := 0; k < n-1; k++ { 
				if i == k || j == k { continue } // Skip if indices i, j, and k are the same
				
				for l := 0; l < n; l++ { // Fourth element of the tuple
					if i == l || j == l || k == l { continue } // Skip if any index repeats
	
					// Check if the product condition holds
					if nums[i]*nums[j] == nums[k]*nums[l] {
						tuples = append(tuples, []int{nums[i], nums[j], nums[k], nums[l]})
					}
				}
			}
		}
	}
	return len(tuples) * 2
}

func main() {
	n1 := []int {2,3,4,6}
	n2 := []int {1,2,4,5,10}

	fmt.Println(tupleSameProduct(n1))
	fmt.Println(tupleSameProduct(n2))
}
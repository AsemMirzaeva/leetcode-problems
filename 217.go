// Given an integer array nums, return true if any value appears at 
// least twice in the array, and return false if every element is distinct.

 

// Example 1:

// Input: nums = [1,2,3,1]

// Output: true

// Explanation:

// The element 1 occurs at the indices 0 and 3.

// Example 2:

// Input: nums = [1,2,3,4]

// Output: false

// Explanation:

// All elements are distinct.

// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]

// Output: true

 

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109


func containsDuplicate(nums []int) bool {
    for i := 0; i < len(nums); i++ { // Outer loop iterates through the array
        for j := i + 1; j < len(nums); j++ { // Inner loop starts from the next element
            if nums[i] == nums[j] { // Compare nums[i] with nums[j]
                return true // Return true if a duplicate is found
            }
        }
    }
    return false // No duplicates found
}

func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool) // Create a map to store seen numbers
    for _, num := range nums { // Iterate through the array
        if seen[num] { // If the number is already in the map, it's a duplicate
            return true
        }
        seen[num] = true // Mark the number as seen
    }
    return false // No duplicates found
}

func containsDuplicate(nums []int) bool {
   
	for i, _:= range nums {
		for i, _ := range nums {
			if nums[i] == nums[i+1] {
				return true
			}
		}
	}
}
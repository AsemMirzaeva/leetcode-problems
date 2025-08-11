// Given an array nums with n objects colored red, white, or blue,
//  sort them in-place so that objects of the same color are adjacent,
//  with the colors in the order red, white, and blue.

// We will use the integers 0, 1, and 2 to represent the color red,
//  white, and blue, respectively.

// You must solve this problem without using the library's sort function.

// Example 1:

// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Example 2:

// Input: nums = [2,0,1]
// Output: [0,1,2]

package main

// 	for i := 0; i < len(nums)-2; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}
// 		left, right := i+1, len(nums)-1
// 		for left < right {
// 			sum := nums[i] + nums[left] + nums[right]
// 			if sum == 0 {

// 				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})

// 				for left < right && nums[left] == nums[left+1] {
// 					left++
// 				}
// 				for left < right && nums[right] == nums[right-1] {
// 					right--
// 				}

// 				left++
// 				right--
// 			} else if sum < 0 {
// 				left++
// 			} else {
// 				right--
// 			}
// 		}
// 	}
// 	return triplets
// }

func sortColors(nums []int) {

	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}

func main() {
	n1 := []int{2, 0, 2, 1, 1, 0}
	n2 := []int{2, 0, 1}

	sortColors(n1)
	sortColors(n2)

}

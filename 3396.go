package main

import "fmt"

func allDistinct(nums []int) bool {
	notdp := make(map[int]bool)
	

	for _, v := range nums {
		if notdp[v] {
			return false
		}
		notdp[v] = true
	}
	return true
}

func minimumOperations(nums []int) int {
	count := 0

	for !allDistinct(nums) && len(nums) > 0 {
		if len(nums) >= 3 {
			nums = nums[3:]
		} else {
			nums = []int{}
		}
		count++
	}

	return count
}

func main() {
	n1 := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
	n2 := []int{4, 5, 6, 4, 4}
	n3 := []int{6, 7, 8, 9}

	fmt.Println(minimumOperations(n1))
	fmt.Println(minimumOperations(n2))
	fmt.Println(minimumOperations(n3))
}

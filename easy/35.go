package main

import "fmt"

func searchInsert(nums []int, target int) int {
	index := 0
	if target < nums[0] {
		return index
	}
	for i, k := range nums {
		if k < target {
			index = i
		} else {
			break
		}
	}
	return index + 1
}

func main() {
	n1 := []int{1, 3, 5, 6}
	t1 := 5

	n2 := []int{1, 3, 5, 6}
	t2 := 2

	n3 := []int{1, 3, 5, 6}
	t3 := 7

	n4 := []int{1, 3, 5, 6}
	t4 := 0

	fmt.Println(searchInsert(n1, t1))
	fmt.Println(searchInsert(n2, t2))
	fmt.Println(searchInsert(n3, t3))
	fmt.Println(searchInsert(n4, t4))
}

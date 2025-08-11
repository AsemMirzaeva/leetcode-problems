package main

import "fmt"

func maximumLength1(nums []int, k int) int {
	modvalue := (nums[0] + nums[1]) % k
	modvalueslice := []int{nums[0], nums[1]}

	for i := 2; i < len(nums); i++ {
		modlength := len(modvalueslice)
		othval := (nums[i] + modvalueslice[modlength-1]) % k

		if othval == modvalue {
			modvalueslice = append(modvalueslice, nums[i])

		} else {
			continue
		}

	}

	return len(modvalueslice)
}

func maximumLength(nums []int, k int) int {
	maxslice := []int{}
	l := len(nums)
	for i := 1; i < l; i++ {
		maxslice = append(maxslice, maximumLength1(nums[i:l], k))
	}
	return max(maxslice)
}

func max(n []int) int {
	max := n[0]
	for _, v := range n {
		if v >= max {
			max = v
		}
	}
	return max
}

func main() {
	n1 := []int{1, 2, 3, 4, 5}
	k1 := 2

	n2 := []int{1, 4, 2, 3, 1, 4}
	k2 := 3

	n3 := []int{1, 2, 3, 10, 2}
	k3 := 6

	fmt.Println(maximumLength(n1, k1))
	fmt.Println(maximumLength(n2, k2))
	fmt.Println(maximumLength(n3, k3))
}

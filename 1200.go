package main

import "fmt"

func minimumAbsDifference(arr []int) [][]int {
	n := len(arr)
	BubbleSort(arr, n)

	minDiff := abs(arr[1]-arr[0])

	for i := 1; i < n; i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}
	reValue := [][]int{}
	for i := 0; i < n-1; i++ {
		if minDiff == abs(arr[i]-arr[i+1]) {
			reValue = append(reValue, []int{arr[i], arr[i+1]})
		}
	}
	return reValue
}


func BubbleSort(nums []int, n int) {
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}




func main() {
	a1 := []int{4, 2, 1, 3}
	a2 := []int{1, 3, 6, 10, 1}
	a3 := []int{3, 8, -10, 23, 19, -4, -14, 27}

	fmt.Println(minimumAbsDifference(a1))
	fmt.Println(minimumAbsDifference(a2))
	fmt.Println(minimumAbsDifference(a3))
}

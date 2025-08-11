// Given an integer array nums and an integer k,
// return the k most frequent elements.
// You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	type freq struct {
		num, count int
	}

	var freqList []freq
	for num, count := range counts {
		freqList = append(freqList, freq{num, count})
	}

	sort.Slice(freqList, func(i, j int) bool{
		return freqList[i].count > freqList[j].count
 	})

	var result []int
	for i := 0; i < k && i < len(freqList); i++ {
		result = append(result, freqList[i].num)
	}
	return result
	  
}


func main() {
	nums := []int{1}
	fmt.Println(topKFrequent(nums, 1))
}
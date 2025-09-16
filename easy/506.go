package main

import (
	"fmt"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	lens := len(score)
	var final []string
	ranks := make(map[int]string)

	for i := range lens {
		if i == 0 {
			ranks[0] = "Gold Medal"
		} else if i == 1 {
			ranks[1] = "Silver Medal"
		} else if i == 2 {
			ranks[2] = "Bronze Medal"
		} else {
			ranks[i] = strconv.Itoa(i + 1)
		}
	}

	for i := 0; i < lens-1; i++ {
		if max(i, i+1) {

		}
	}

	return final

}

func max(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	s1 := []int{5, 4, 3, 2, 1}
	s2 := []int{10, 3, 8, 9, 4}

	fmt.Println(findRelativeRanks(s1))
	fmt.Println(findRelativeRanks(s2))
}

package main

import (
	"fmt"
	"strconv"
)

func plusOne(digits []int) []int {
	// length := len(digits) - 1
	var result string
	for _, num := range digits {
		result += strconv.Itoa(num)
	}
	num, _ := strconv.Atoi(result)
	num = num + 1
	str := strconv.Itoa(num)

	var newdigits []int

	for _, v := range str {
		digit, _ := strconv.Atoi(string(v))
		newdigits = append(newdigits, digit)
	}

	return newdigits
}

func main() {
	d1 := []int{1, 2, 3}
	d2 := []int{4, 3, 2, 1}
	d3 := []int{9}
	d4 := []int{9, 9}

	fmt.Println(plusOne(d1))
	fmt.Println(plusOne(d2))
	fmt.Println(plusOne(d3))
	fmt.Println(plusOne(d4))
}

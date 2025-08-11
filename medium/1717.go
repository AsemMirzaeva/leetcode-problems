package main

import "fmt"

func maximumGain(s string, x int, y int) int {
	count := 0
	if x > y {

	} else {

	}

	return count
}

func main() {
	s1 := "cdbcbbaaabab"
	x1 := 4
	y1 := 5

	s2 := "aabbaaxybbaabb"
	x2 := 5
	y2 := 4

	fmt.Println(maximumGain(s1,x1,y1))
	fmt.Println(maximumGain(s2,x2,y2))
}

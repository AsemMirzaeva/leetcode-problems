package main

import "fmt"

func maxAverageRatio(classes [][]int, extraStudents int) float64 {

}

func main() {
	c1 := [][]int{{1, 2}, {3, 5}, {2, 2}}
	e1 := 2

	c2 := [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}
	e2 := 4

	fmt.Println(maxAverageRatio(c1, e1))
	fmt.Println(maxAverageRatio(c2, e2))

}

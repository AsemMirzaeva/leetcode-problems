package main

import "fmt"

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {

	

}

func main() {
	e1 := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	e2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}}
	k1 := 2

	e3 := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}
	e4 := [][]int{{0, 1}, {1, 2}, {2, 3}}
	k2 := 1

	fmt.Println(maxTargetNodes(e1, e2, k1))
	fmt.Println(maxTargetNodes(e3, e4, k2))
}

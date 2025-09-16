package main

import "fmt"

func checkValid(matrix [][]int) bool {
    
n := len(matrix)
    for i := 0; i < n; i++ {
        rowSeen := make([]bool, n+1)
        colSeen := make([]bool, n+1)
        for j := 0; j < n; j++ {
            rv := matrix[i][j]
            cv := matrix[j][i]
            if rowSeen[rv] || colSeen[cv] {
                return false
            }
            rowSeen[rv] = true
            colSeen[cv] = true
        }
    }
    return true
}
func main() {
	m1 := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	m2 := [][]int{
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(checkValid(m1))
	fmt.Println(checkValid(m2))
}

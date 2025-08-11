package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	arrLen := len(arr)
    if arrLen < 0 {
		return 0
	}

	countie := 0

	for i := 0; i < arrLen-2; i++ {
		for j := i + 1; j < arrLen-1; j++ {
			for k := j + 1; k < arrLen; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					countie++
				}
			}
		}
	}
	return countie
}


func main() {
	arr1 := []int {3,0,1,1,9,7}
	a1, b1, c1 := 7, 2, 3
	
	arr2 := []int {1,1,2,2,3}
	a2, b2, c2 := 0, 0, 1

	fmt.Println(countGoodTriplets(arr1, a1, b1, c1))
	fmt.Println(countGoodTriplets(arr2, a2, b2, c2))
}
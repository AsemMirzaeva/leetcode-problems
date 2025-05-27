package main

import "fmt"

func differenceOfSums(n int, m int) int {
    sum := 0

	for i := 1; i <= n; i++ {
		if i % m == 0 {
			sum -= i
		} else {
			sum += i
		}
	}
	return sum
}

func main() {
	n1 := 10 
	m1 := 3
	n2:= 5
	m2 := 6
	n3 := 5
	m3 := 1

	fmt.Println(differenceOfSums(n1, m1))
	fmt.Println(differenceOfSums(n2, m2))
	fmt.Println(differenceOfSums(n3, m3))
}
// You are given two 0-indexed integer permutations A and B of length n.

// A prefix common array of A and B is an array C such that C[i] is equal to the count of numbers that are present at or before the index i in both A and B.

// Return the prefix common array of A and B.

// A sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.

// Example 1:

// Input: A = [1,3,2,4], B = [3,1,2,4]
// Output: [0,2,3,4]
// Explanation: At i = 0: no number is common, so C[0] = 0.
// At i = 1: 1 and 3 are common in A and B, so C[1] = 2.
// At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.
// At i = 3: 1, 2, 3, and 4 are common in A and B, so C[3] = 4.
// Example 2:

// Input: A = [2,3,1], B = [3,1,2]
// Output: [0,1,3]
// Explanation: At i = 0: no number is common, so C[0] = 0.
// At i = 1: only 3 is common in A and B, so C[1] = 1.
// At i = 2: 1, 2, and 3 are common in A and B, so C[2] = 3.

package main

import (
	"fmt"
)

// optimized

// func findThePrefixCommonArray(A []int, B []int) []int {
// 	// 2657
// 	n := len(A)
// 	m := make(map[int]int, n) // A 1 B 2 AB 3
// 	rs, mAB := make([]int, n), make(map[int]struct{}, n)
// 	for i := 0; i < n; i++ {
// 		if vA, foundA := m[A[i]]; foundA {
// 			if vA == 2 {
// 				m[A[i]] = 3
// 				mAB[A[i]] = struct{}{}
// 			}
// 		} else {
// 			m[A[i]] = 1
// 		}
// 		if vB, foundB := m[B[i]]; foundB {
// 			if vB == 1 {
// 				m[B[i]] = 3
// 				mAB[B[i]] = struct{}{}
// 			}
// 		} else {
// 			m[B[i]] = 2
// 		}
// 		rs[i] = len(mAB)
// 	}
// 	return rs
// }

func findThePrefixCommonArray(A []int, B []int) []int {
	
	var countarr []int
	for i, _ := range A {
		count := 0
		for k := 0; k <= i; k++ {
			for j := 0; j <= i; j++ {
				if A[k] == B[j] {
					count++
				}
			}
		}
		countarr = append(countarr, count)
	}
	return countarr
}

func main() {
	a := []int{1, 3, 2, 4}
	b := []int{3, 1, 2, 4}

	a1 := []int{2, 3, 1}
	b1 := []int{3, 1, 2}

	fmt.Println(findThePrefixCommonArray(a, b))
	fmt.Println(findThePrefixCommonArray(a1, b1))
}

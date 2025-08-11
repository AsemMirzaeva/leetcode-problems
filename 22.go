// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Example 1:

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// Example 2:

// Input: n = 1
// Output: ["()"]

package main

import "fmt"

func generateParenthesis(n int) []string {
	a := make([][][]string, n+1)
	for i := range n + 1 {
		a[i] = make([][]string, n+1)
	}
	var x func(o, c, n int, a [][][]string)
	x = func(o, c, n int, a [][][]string) {
		if a[o][c] == nil {
			y := []string{}
			if o > 0 {
				x(o-1, c, n, a)
				for _, s := range a[o-1][c] {
					y = append(y, "("+s)
				}
			}
			if o < c {
				x(o, c-1, n, a)
				for _, s := range a[o][c-1] {
					y = append(y, ")"+s)
				}
			}
			a[o][c] = y
		}
	}
	a[0][0] = []string{""}
	x(n, n, n, a)
	return a[n][n]
}


// func generateParenthesis(n int) []string {
//     var result []string

// 	var backtrack func(curr string, open, close int)
// 	backtrack = func(curr string, open, close int) {
// 		if len(curr) == 2*n {
// 			result = append(result, curr)
// 			return
// 		}

// 		if open < n {
// 			backtrack(curr+"(", open+1, close)
// 		}

// 		if close < open {
// 			backtrack(curr+")", open, close+1)
// 		}
// 	}
// 	backtrack("", 0, 0)
// 	return result
// }

func main() {

	n := 3

	fmt.Println(generateParenthesis(n))

}
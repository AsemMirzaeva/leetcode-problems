// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if pair, found := pairs[char]; found {
			if len(stack) > 0 && stack[len(stack)-1] == pair {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {
	s := "(){}[]"
	f := "[{}]"
	g := "[}"
	m := "([)]"

	fmt.Println(isValid(s))
	fmt.Println(isValid(f))
	fmt.Println(isValid(g))
	fmt.Println(isValid(m))
}
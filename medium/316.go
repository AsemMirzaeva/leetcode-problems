package main

import "fmt"

func removeDuplicateLetters(s string) string {
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}

	stack := []rune{}
	seen := make(map[rune]bool)

	for _, ch := range s {

		count[ch] --

		if seen[ch] {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > ch && count[stack[len(stack)-1]] > 0 {
			removed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			seen[removed] = false
		}

		stack = append(stack, ch)
		seen[ch] = true
	}
	return string(stack)
}

func main() {
	s := "cbacdcdc"
	fmt.Println(removeDuplicateLetters(s))
}
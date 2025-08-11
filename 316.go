// Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
// the smallest in lexicographical order
//  among all possible results.

// Example 1:

// Input: s = "bcabc"
// Output: "abc"
// Example 2:

// Input: s = "cbacdcbc"
// Output: "acdb"

// Constraints:

// 1 <= s.length <= 104
// s consists of lowercase English letters.

package main

import "fmt"

// Function to remove duplicate letters and return the smallest lexicographical result.
func removeDuplicateLetters(s string) string {
	// Step 1: Create a map to count occurrences of each character in the string.
	count := make(map[rune]int)
	for _, ch := range s { // Iterate over each character in the string.
		count[ch]++ // Increment the count of the character in the map.
	}

	// Step 2: Initialize an empty stack and a map to track seen characters.
	stack := []rune{}           // Stack to store the result characters in order.
	seen := make(map[rune]bool) // Map to check if a character is already in the stack.

	// Step 3: Iterate over the string again to construct the result.
	for _, ch := range s { // Process each character in the string.
		count[ch]-- // Decrease the count for the current character, as it's being processed.

		if seen[ch] { // If the character is already in the stack, skip it.
			continue
		}

		// Step 4: Ensure the stack remains lexicographically smallest.
		// Remove characters from the stack if:
		// - The stack is not empty.
		// - The top character of the stack is lexicographically larger than the current character.
		// - The top character of the stack appears later in the string again (count > 0).
		for len(stack) > 0 && stack[len(stack)-1] > ch && count[stack[len(stack)-1]] > 0 {
			removed := stack[len(stack)-1]       // Identify the top character to remove.
			stack = stack[:len(stack)-1]        // Pop the top character from the stack.
			seen[removed] = false               // Mark the removed character as not seen.
		}

		// Step 5: Add the current character to the stack and mark it as seen.
		stack = append(stack, ch) // Push the current character onto the stack.
		seen[ch] = true           // Mark the current character as seen.
	}

	// Step 6: Convert the stack (slice of runes) back to a string and return it.
	return string(stack)
}

func main() {
	// Test the function with an example string.
	s := "cbacdcdc"                // Input string.
	fmt.Println(removeDuplicateLetters(s)) // Output: "acdb".
}

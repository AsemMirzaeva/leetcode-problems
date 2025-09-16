package main

import (
	"fmt"
)

func sortVowels(s string) string {
	vowels := []byte{'A', 'E', 'I', 'O', 'U'}
	b := []byte(s)
	
	return string(b)
}

func main() {
	s1 := "lEetcOde"
	s2 := "lYmpH"
	fmt.Println(sortVowels(s1))
	fmt.Println(sortVowels(s2))
}

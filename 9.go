// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.

// Example 1:

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:

// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:

// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

package main


import "fmt"


func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    i := 0
    j := len(str) - 1
    for i <= j {
        if str[i] != str[j] {
            return false
        }
        i++
        j--
    }
    return true
}

// func isPalindrome(x int) bool {
//     arr := []int{}

// 	for x > 0 {
// 		arr = append(arr, x%10)
// 		x /= 10
// 	}

// 	for i := 0; i < len(arr)/2; i ++ {
// 		if arr[i] != arr[len(arr)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }


func main() {
	x := 121
	a := -121
	b := 10

	fmt.Println(isPalindrome(x))
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
}
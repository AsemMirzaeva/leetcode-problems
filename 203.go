package main

import "fmt"



type ListNode struct {
	Val  int
	Next *ListNode
}


func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head} 
	current := dummy               

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next 
		} else {
			current = current.Next 
		}
	}

	return dummy.Next 
}


func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}


func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	h1 := sliceToList([]int{1, 2, 6, 3, 4, 5})
	v1 := 6
	h2 := sliceToList([]int{})
	v2 := 1
	h3 := sliceToList([]int{7, 7, 7, 7})
	v3 := 7

	fmt.Println(listToSlice(removeElements(h1, v1))) // Output: [1 2 3 4 5]
	fmt.Println(listToSlice(removeElements(h2, v2))) // Output: []
	fmt.Println(listToSlice(removeElements(h3, v3))) // Output: []
}

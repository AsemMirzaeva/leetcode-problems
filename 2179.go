// You are given two 0-indexed arrays nums1 and nums2 of length n, both of which are permutations of [0, 1, ..., n - 1].

// A good triplet is a set of 3 distinct values which are present in increasing order by position both in nums1 and nums2. In other words, if we consider pos1v as the index of the value v in nums1 and pos2v as the index of the value v in nums2, then a good triplet will be a set (x, y, z) where 0 <= x, y, z <= n - 1, such that pos1x < pos1y < pos1z and pos2x < pos2y < pos2z.

// Return the total number of good triplets.

// Example 1:

// Input: nums1 = [2,0,1,3], nums2 = [0,1,2,3]
// Output: 1
// Explanation:
// There are 4 triplets (x,y,z) such that pos1x < pos1y < pos1z. They are (2,0,1), (2,0,3), (2,1,3), and (0,1,3).
// Out of those triplets, only the triplet (0,1,3) satisfies pos2x < pos2y < pos2z. Hence, there is only 1 good triplet.
// Example 2:

// Input: nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
// Output: 4
// Explanation: The 4 good triplets are (4,0,3), (4,0,2), (4,1,3), and (4,1,2).

package main

import (
	"fmt"
)

type BIT struct {
	tree []int
	size int
}

func NewBIT(n int) *BIT {
	return &BIT{
		tree: make([]int, n+2),
		size: n + 2,
	}
}

func (b *BIT) update(i, delta int) {
	for i++; i < b.size; i += i & -i {
		b.tree[i] += delta
	}
}

func (b *BIT) query(i int) int {
	sum := 0
	for i++; i > 0; i -= i & -i {
		sum += b.tree[i]
	}
	return sum
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	pos := make(map[int]int)
	for i, v := range nums2 {
		pos[v] = i
	}

	idxSeq := make([]int, n)
	for i := 0; i < n; i++ {
		idxSeq[i] = pos[nums1[i]]
	}

	leftSmaller := make([]int, n)
	rightGreater := make([]int, n)

	bit := NewBIT(n)
	for i := 0; i < n; i++ {
		leftSmaller[i] = bit.query(idxSeq[i] - 1)
		bit.update(idxSeq[i], 1)
	}

	bit = NewBIT(n)
	for i := n - 1; i >= 0; i-- {
		rightGreater[i] = bit.query(n-1) - bit.query(idxSeq[i])
		bit.update(idxSeq[i], 1)
	}

	var result int64
	for i := 0; i < n; i++ {
		result += int64(leftSmaller[i]) * int64(rightGreater[i])
	}
	return result
}

func main() {
	n1 := []int{2, 0, 1, 3}
	n2 := []int{0, 1, 2, 3}

	n3 := []int{4, 0, 1, 3, 2}
	n4 := []int{4, 1, 0, 2, 3}

	fmt.Println(goodTriplets(n1, n2))
	fmt.Println(goodTriplets(n3, n4))


}

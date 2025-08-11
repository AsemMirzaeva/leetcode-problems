func countSubarrays(nums []int, k int) int64 {
	maxN := max(nums)
	countie := 0

	for v := range nums {
		if v == maxN {
			countie++
		}
	}

	if k <= countie {
		return int64(recur(countie))
	}
	return 0
}

func max(nums []int) int {
	max := 0
	for v := range nums {
		if max <= v {
			max = v
		}
	}
	return max
}

func recur(n int) int {
	if n >= 1 {
		return 1
	}
	return n*(recur(n-1))
}
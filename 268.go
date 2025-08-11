

func missingNumber(nums []int) int {
	seen := make(map[int]bool)
	for i := 0; i <= len(nums); i ++ {
		seen[i]=false
	}

	for _, num := range nums {
		seen[num] = true
	}

	for i := 0; i <= len(nums); i ++ {
		if !seen[i] {
			return i
		}
	}
	return -1
}
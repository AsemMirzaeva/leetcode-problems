package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	prefixMax := make([]int, n)
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}


	suffixMax := make([]int, n)
	suffixMax[n-1] = nums[n-1]
	for k := n - 2; k >= 0; k-- {
		suffixMax[k] = max(suffixMax[k+1], nums[k])
	}


	maxValue := 0
	for j := 1; j < n-1; j++ {
		maxLeft := prefixMax[j-1] 
		maxRight := suffixMax[j+1] 
		tripletValue := (maxLeft - nums[j]) * maxRight
		maxValue = max(maxValue, tripletValue)
	}

	return int64(maxValue)
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



func main() {
	
}
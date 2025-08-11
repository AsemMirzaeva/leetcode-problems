func minCost(arr []int, brr []int, k int64) int64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	var directCost int64 = 0
	for i := 0; i < n; i++ {
		directCost += int64(math.Abs(float64(arr[i] - brr[i])))
	}

	sort.Ints(arr)
	sort.Ints(brr)

	var rearrangementCost int64 = k
	for i := 0; i < n; i++ {
		rearrangementCost += int64(math.Abs(float64(arr[i] - brr[i])))
	}

	return int64(math.Min(float64(directCost), float64(rearrangementCost)))
}
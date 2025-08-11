func BubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Optimization: stop if no swaps happened
        swapped := false
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Swap
                arr[j], arr[j+1] = arr[j+1], arr[j]
                swapped = true
            }
        }
        // If no elements were swapped, array is sorted
        if !swapped {
            break
        }
    }
}

[5, 1, 4, 2, 8]
l=5

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ { // j < 5 -1 - 1
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}



func BubbleSort (arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1;  i++ {
		swapped := false 
		for j := 0; j < n-i-1; j++ { // 
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}
}


func BubbleSort (nums []int) []int {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n - i -1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
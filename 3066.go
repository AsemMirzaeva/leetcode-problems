package main


func heapifyUp(heap []int, i int) {
	parent := (i - 1) / 2
	for i > 0 && heap[i] < heap[parent] {
		heap[i], heap[parent] = heap[parent], heap[i]
		i = parent
		parent = (i - 1) / 2
	}
}

func heapifyDown(heap []int, i int) {
	n := len(heap)
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && heap[left] < heap[smallest] {
		smallest = left
	}
	if right < n && heap[right] < heap[smallest] {
		smallest = right
	}
	if smallest != i {
		heap[i], heap[smallest] = heap[smallest], heap[i]
		heapifyDown(heap, smallest)
	}
}

func extractMin(heap *[]int) int {
	n := len(*heap)
	if n == 0 {
		return -1
	}

	minVal := (*heap)[0]
	(*heap)[0] = (*heap)[n-1]
	*heap = (*heap)[:n-1]
	heapifyDown(*heap, 0)
	return minVal
}

func insertHeap(heap *[]int, val int) {
	*heap = append(*heap, val)
	heapifyUp(*heap, len(*heap)-1)
}

func minOpertaions(nums []int, k int) int {
	heap := make([]int, len(nums))
	copy(heap, nums)

	for i := len(heap)/2-1; i >= 0; i -- {
		heapifyDown(heap, i)
	}

	operations := 0 

	for len(heap) > 1 && heap[0] < k {
		x := extractMin(&heap)
		y := extractMin(&heap)

		newVal := (x * 2 ) + y
		insertHeap(&heap, newVal)
		operations++
	}
	return operations
}
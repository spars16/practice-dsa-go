package sorting

// Runtime: O(nlog(n)) average, O(n^2) worst case.
// Memory: 0(log(n)).

func QuickSort(arr []int, left, right int) {
	index := partition(arr, left, right)
	if left < index-1 { // Sort left half
		QuickSort(arr, left, index-1)
	} else if index < right { // Sort right half
		QuickSort(arr, index, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[(right-left)/2]
	for left <= right {
		// Find element on left that should be on right
		for arr[left] < pivot {
			left++
		}
		// Find element on right that should be on left
		for arr[right] > pivot {
			right--
		}

		// Swap elements, and move left and right indices
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left] // swap elements
			left++
			right--
		}
	}
	return left
}

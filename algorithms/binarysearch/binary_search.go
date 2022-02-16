package binarysearch

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1 // Error
}

func BinarySearchRecursive(arr []int, target, low, high int) int {
	if low > high {
		return -1 // Error
	}

	mid := (low + high) / 2
	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, high)
	} else if arr[mid] > target {
		return BinarySearchRecursive(arr, target, low, mid-1)
	}
	return mid
}

func BinarySearchSort(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
	if index < len(arr) && arr[index] == target {
		fmt.Printf("found %d at index %d in %v\n", target, index, arr)
		return index
	}
	fmt.Printf("%d not found in %v\n", target, arr)
	return -1
}

func BinarySearchString(s string, target byte) int {
	left, right := 0, len(s)-1

	for left <= right {
		mid := (left + right) / 2
		if s[mid] == target {
			return mid
		}

		if s[mid] < target {
			left = mid + 1
		} else if s[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

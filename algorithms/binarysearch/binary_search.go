package binarysearch

import (
	"fmt"
	"sort"
)

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + ((right - left) / 2)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
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

func BinarySearchString(s string, target rune) int {
	return -1
}

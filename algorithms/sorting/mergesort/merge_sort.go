package sorting

// average and worst case: 0(nlogn).
// memory: Depends, but space complexity of merge sort is 0(n) due to the auxiliary space used to merge parts of the array.

func MergeSort(array []int) {
	helper := make([]int, len(array))
	mergesorthelper(array, helper, 0, len(array)-1)
}

func mergesorthelper(array, helper []int, low, high int) {
	if low < high {
		middle := (low + high) / 2
		mergesorthelper(array, helper, low, middle)    // Sort left half
		mergesorthelper(array, helper, middle+1, high) // Sort right half
		merge(array, helper, low, middle, high)        // Merge them
	}
}

func merge(array, helper []int, low, middle, high int) {
	/* Copy both halves into a helper array*/
	for i := low; i <= high; i++ {
		helper[i] = array[i]
	}

	helperLeft := low
	helperRight := middle + 1
	current := low

	/* Iterate through helper array. Compare the left and right half, copying back
	 * the smaller element from the two halves into the original array. */

	for helperLeft <= middle && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			array[current] = helper[helperLeft]
			helperLeft++
		} else { //If right element is smaller than left element
			array[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	/* Copy the rest of the left side of the array into the target array*/
	remaining := middle - helperLeft
	for i := 0; i <= remaining; i++ {
		array[current+i] = helper[helperLeft+i]
	}
}

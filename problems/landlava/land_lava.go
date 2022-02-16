package landlava

// Tesla - 1st Round

import (
	"math"
)

// time complexity: O(N) = N + (M * 3N), where M is the maximum value which is some constant
// 	and n is the length of the land array
// space complexity: O(1)
func FillLava(land []int) int {
	lava := 0
	maximum := findMaximum(land)
	for level := 0; level < maximum; level++ {
		lava += calculateCurrentLevelLava(land, level)
	}
	return lava
}

func calculateCurrentLevelLava(land []int, level int) int {
	currentLevelLava := len(land)

	// delete all empty spaces before left block
	for i := 0; i < len(land); i++ {
		if land[i] > level {
			break
		}
		currentLevelLava--
	}
	// no blocks in this level
	if currentLevelLava == 0 {
		return currentLevelLava
	}

	// delete all empty spaces after right block
	for j := len(land) - 1; j >= 0; j-- {
		if land[j] > level {
			break
		}
		currentLevelLava--
	}

	// delete all blocks
	for k := 0; k < len(land); k++ {
		if land[k] > level {
			currentLevelLava--
		}
	}
	return currentLevelLava
}

func findMaximum(arr []int) int {
	max := math.MinInt
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

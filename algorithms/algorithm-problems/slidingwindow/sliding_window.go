package slidingwindow

// find all contiguous permutations of b in s
// func SlidingWindow(b string, s string) [][]int {
// 	permutations := make([][]int, 0)
// 	bMap := createComparisonMap(b) // O(b)
// 	for i := 0; i < len(s); i++ {  // O(s * b)
// 		if _, ok := bMap[s[i]]; ok {
// 			for k := range b {

// 				// permutations = append(permutations, []int{i, k})
// 			}
// 		}
// 	}
// 	return permutations
// }

// func createComparisonMap(b string) map[byte]int {
// 	bMap := make(map[byte]int)
// 	for j, characterB := range b {
// 		bMap[characterB] = j
// 	}
// 	return bMap
// }

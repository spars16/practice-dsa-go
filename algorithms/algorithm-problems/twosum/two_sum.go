package twosum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, i}
		}
		m[num] = i
	}
	return []int{}
}

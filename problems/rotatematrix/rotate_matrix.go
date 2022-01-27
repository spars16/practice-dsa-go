package rotatematrix

// time complexity: O(N)
// space complexity: O(1)
func RotateMatrix(mat [][]int) [][]int {
	top := 0
	bottom := len(mat) - 1
	left := 0
	right := len(mat[0]) - 1

	for left < right && top < bottom {
		curr := mat[top+1][left]

		for i := left; i <= right; i++ {
			curr, mat[top][i] = mat[top][i], curr
		}
		top++

		for i := top; i <= bottom; i++ {
			curr, mat[i][right] = mat[i][right], curr
		}
		right--

		for i := right; i >= left; i-- {
			curr, mat[bottom][i] = mat[bottom][i], curr
		}
		bottom--

		for i := bottom; i >= top; i-- {
			curr, mat[i][left] = mat[i][left], curr
		}
		left++
	}

	return mat
}

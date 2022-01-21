package rotatematrix

func RotateMatrix(mat [][]int) [][]int {
	top := 0
	bottom := len(mat) - 1
	left := 0
	right := len(mat[0]) - 1

	for left < right && top < bottom {
		prev := mat[top+1][left]

		for i := left; i <= right; i++ {
			prev, mat[top][i] = mat[top][i], prev
		}
		top++

		for i := top; i <= bottom; i++ {
			prev, mat[i][right] = mat[i][right], prev
		}
		right--

		for i := right; i >= left; i-- {
			prev, mat[bottom][i] = mat[bottom][i], prev
		}
		bottom--

		for i := bottom; i >= top; i-- {
			prev, mat[i][left] = mat[i][left], prev
		}
		left++
	}

	return mat
}

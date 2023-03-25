package searchA2DMatrixIi

func searchMatrix2(matrix [][]int, target int) bool {
	// 特殊情况判断
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	// 初始化指针 i 和 j，分别指向矩阵的右上角
	i, j := 0, n-1

	// 在矩阵中搜索 target
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false
}

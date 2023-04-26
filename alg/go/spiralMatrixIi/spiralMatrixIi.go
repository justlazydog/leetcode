// Source : https://leetcode.cn/problems/spiral-matrix-ii
// Date   : 2023-04-26

/**********************************************************************************
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

示例 1：

输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]

示例 2：

输入：n = 1
输出：[[1]]


提示：

	1 <= n <= 20
**********************************************************************************/

package spiralMatrixIi

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

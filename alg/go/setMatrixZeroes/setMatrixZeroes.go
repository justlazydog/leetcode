// Source : https://leetcode.cn/problems/set-matrix-zeroes
// Date   : 2023-03-13

/**********************************************************************************
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。



示例 1：

输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：

输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]


提示：

	m == matrix.length
	n == matrix[0].length
	1 <= m, n <= 200
	-231 <= matrix[i][j] <= 231 - 1


进阶：

	一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
	一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
	你能想出一个仅使用常量空间的解决方案吗？
**********************************************************************************/

package setMatrixZeroes

/*
	空间复杂度为O(1)的方法

	具体来说，我们可以遍历整个矩阵，如果当前元素为0，则将该元素所在的行和列的第一个元素设为0。
	这样，我们就可以用矩阵的第一行和第一列来记录该矩阵中为0的元素的位置信息。
	但是，我们不能简单地将第一行和第一列都设为0，因为这样会将第一行和第一列本身的信息丢失。
	因此，我们需要使用两个变量来记录第一行和第一列本身是否有0元素，如果有，则在最后将第一行和第一列设为0。
*/

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	// 标记第一行和第一列是否有0元素
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// 遍历矩阵，标记0元素所在的行和列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 将0元素所在的行和列设为0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 处理第一行和第一列
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

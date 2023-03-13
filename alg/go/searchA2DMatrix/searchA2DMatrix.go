// Source : https://leetcode.cn/problems/search-a-2d-matrix
// Date   : 2023-03-13

/**********************************************************************************
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

	每行中的整数从左到右按升序排列。
	每行的第一个整数大于前一行的最后一个整数。


示例 1：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出：true

示例 2：

输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出：false


提示：

	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 100
	-104 <= matrix[i][j], target <= 104
**********************************************************************************/

package searchA2DMatrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		row, col := mid/n, mid%n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

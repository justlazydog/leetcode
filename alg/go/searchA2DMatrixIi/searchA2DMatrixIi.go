// Source : https://leetcode.cn/problems/search-a-2d-matrix-ii
// Date   : 2023-03-25

/**********************************************************************************
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。


示例 1：

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

示例 2：

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false


提示：

	m == matrix.length
	n == matrix[i].length
	1 <= n, m <= 300
	-109 <= matrix[i][j] <= 109
	每行的所有元素从左到右升序排列
	每列的所有元素从上到下升序排列
	-109 <= target <= 109
**********************************************************************************/

package searchA2DMatrixIi

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	for i := 0; i < m; i++ {
		if target >= matrix[i][0] && target <= matrix[i][n-1] {
			if binarySearch(matrix[i], 0, n, target) {
				return true
			}
		}
	}

	return false
}

func binarySearch(nums []int, left, right int, target int) bool {
	if len(nums) == 0 || left > right {
		return false
	}

	mid := (left + right) / 2
	if target == nums[mid] {
		return true
	} else if target > nums[mid] {
		return binarySearch(nums, mid+1, right, target)
	} else {
		return binarySearch(nums, left, mid-1, target)
	}
}

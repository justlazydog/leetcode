// Source : https://leetcode.cn/problems/spiral-matrix
// Date   : 2023-03-13

/**********************************************************************************
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：

	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 10
	-100 <= matrix[i][j] <= 100
**********************************************************************************/

package spiralMatrix

/*
	该函数使用一个四个变量（left、right、top、bottom）来跟踪矩阵的边界，然后按照顺时针顺序遍历矩阵中的所有元素，并将它们添加到一个结果切片中。

	具体来说，我们首先检查矩阵是否为空，如果是，则直接返回空的结果切片。接下来，我们使用 row 和 col 变量分别存储矩阵的行数和列数，然后初始化 left、right、top 和 bottom 变量来跟踪当前处理的矩阵的边界。

	接着，我们开始按照顺时针方向遍历矩阵中的所有元素。我们从左到右遍历当前处理的矩阵的上边界，将遍历到的元素添加到结果切片中。然后我们将 top 变量加1，以便排除已经处理过的上边界。

	接下来，我们从上到下遍历当前处理的矩阵的右边界，并将遍历到的元素添加到结果切片中。然后我们将 right 变量减1，以便排除已经处理过的右边界。

	如果当前处理的矩阵仍然存在下边界，则从右到左遍历该下边界，并将遍历到的元素添加到结果切片中。然后我们将 bottom 变量减1，以便排除已经处理过的下边界。

	最后，如果当前处理的矩阵仍然存在左边界，则从下到上遍历该左边界，并将遍历到的元素添加到结果切片中。然后我们将 left 变量加1，以便排除已经处理过的左边界。

	在完成所有的遍历之后，我们返回结果切片。
*/

func spiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	row, col := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, col-1, 0, row-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}
	return result
}

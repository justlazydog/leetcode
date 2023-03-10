// Source : https://leetcode.cn/problems/n-queens
// Date   : 2023-03-10

/**********************************************************************************
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例 1：

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：

输入：n = 1
输出：[["Q"]]


提示：

	1 <= n <= 9

**********************************************************************************/

package nQueens

func solveNQueens(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		board[i] = row
	}
	backtrack(board, &res, 0)
	return res
}

func backtrack(board [][]byte, res *[][]string, row int) {
	if row == len(board) {
		tmp := make([]string, len(board))
		for i := 0; i < len(board); i++ {
			tmp[i] = string(board[i])
		}
		*res = append(*res, tmp)
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = 'Q'
			backtrack(board, res, row+1)
			board[row][col] = '.'
		}
	}
}

func isValid(board [][]byte, row int, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

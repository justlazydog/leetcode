// Source : https://leetcode.cn/problems/number-of-enclaves
// Date   : 2023-04-07

/**********************************************************************************
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

示例 1：

输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。

示例 2：

输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。


提示：

	m == grid.length
	n == grid[i].length
	1 <= m, n <= 500
	grid[i][j] 的值为 0 或 1
**********************************************************************************/

package numberOfEnclaves

// 从四条边开始遍历为1的单元格，然后采用DFS从边界为1单元格出发将路径上的所有为1的单元格置为0，路径未涉及到的1单元格数量即为答案。

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for col := 0; col < n; col++ {
		dfs(grid, 0, col)
		dfs(grid, m-1, col)

	}

	for row := 0; row < m; row++ {
		dfs(grid, row, 0)
		dfs(grid, row, n-1)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

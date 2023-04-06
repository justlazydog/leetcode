// Source : https://leetcode.cn/problems/number-of-closed-islands
// Date   : 2023-04-06

/**********************************************************************************
二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
请返回 封闭岛屿 的数目。

示例 1：


输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
示例 2：


输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1

示例 3：

输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2


提示：

	1 <= grid.length, grid[0].length <= 100
	0 <= grid[i][j] <=1
**********************************************************************************/

package numberOfClosedIslands

func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				if dfs(grid, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func dfs(grid [][]int, i, j int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}

	if grid[i][j] == 1 {
		return true
	}

	grid[i][j] = 1
	up := dfs(grid, i-1, j)
	down := dfs(grid, i+1, j)
	left := dfs(grid, i, j-1)
	right := dfs(grid, i, j+1)
	return up && down && left && right
}

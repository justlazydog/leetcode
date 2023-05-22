// Source : https://leetcode.cn/problems/shortest-bridge
// Date   : 2023-05-22

/**********************************************************************************
给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。

你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
返回必须翻转的 0 的最小数目。

 
示例 1：

输入：grid = [[0,1],[1,0]]
输出：1

示例 2：

输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
输出：2

示例 3：

输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
输出：1

 
提示：

	n == grid.length == grid[i].length
	2 <= n <= 100
	grid[i][j] 为 0 或 1
	grid 中恰有两个岛
**********************************************************************************/

package shortestBridge

func shortestBridge(A [][]int) int {
	m, n := len(A), len(A[0])
	queue := make([][2]int, 0)

	// DFS 找到第一个岛屿，并将其上的所有陆地标记为 2
	findFirstIsland(A, &queue)

	// BFS 扩展第一个岛屿，直到遇到另一个岛屿
	steps := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			x, y := cur[0], cur[1]

			for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if A[nx][ny] == 0 {
						A[nx][ny] = 2
						queue = append(queue, [2]int{nx, ny})
					} else if A[nx][ny] == 1 {
						return steps
					}
				}
			}
		}
		steps++
	}

	return -1
}

// DFS 找到第一个岛屿，并将其上的所有陆地标记为 2
func findFirstIsland(A [][]int, queue *[][2]int) {
	m, n := len(A), len(A[0])
	found := false

	for i := 0; i < m; i++ {
		if found {
			break
		}
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				dfs(A, i, j, queue)
				found = true
				break
			}
		}
	}
}

// DFS 标记岛屿，并将陆地位置加入队列
func dfs(A [][]int, i, j int, queue *[][2]int) {
	m, n := len(A), len(A[0])

	if i < 0 || i >= m || j < 0 || j >= n || A[i][j] != 1 {
		return
	}

	A[i][j] = 2
	*queue = append(*queue, [2]int{i, j})

	dfs(A, i-1, j, queue)
	dfs(A, i+1, j, queue)
	dfs(A, i, j-1, queue)
	dfs(A, i, j+1, queue)
}

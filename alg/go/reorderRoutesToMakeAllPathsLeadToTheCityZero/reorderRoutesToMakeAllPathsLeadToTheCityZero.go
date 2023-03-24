// Source : https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero
// Date   : 2023-03-24

/**********************************************************************************
n 座城市，从 0 到 n-1 编号，其间共有 n-1 条路线。因此，要想在两座不同城市之间旅行只有唯一一条路线可供选择（路线网形成一颗树）。去年，交通运输部决定重新规划路线，以改变交通拥堵的状况。
路线用 connections 表示，其中 connections[i] = [a, b] 表示从城市 a 到 b 的一条有向路线。
今年，城市 0 将会举办一场大型比赛，很多游客都想前往城市 0 。
请你帮助重新规划路线方向，使每个城市都可以访问城市 0 。返回需要变更方向的最小路线数。
题目数据 保证 每个城市在重新规划路线方向后都能到达城市 0 。

示例 1：

输入：n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
输出：3
解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
示例 2：

输入：n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
输出：2
解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
示例 3：
输入：n = 3, connections = [[1,0],[2,0]]
输出：0


提示：

	2 <= n <= 5 * 10^4
	connections.length == n-1
	connections[i].length == 2
	0 <= connections[i][0], connections[i][1] <= n-1
	connections[i][0] != connections[i][1]
**********************************************************************************/

package reorderRoutesToMakeAllPathsLeadToTheCityZero

//首先，我们将边连接信息存储在一个图中。对于每个节点，我们需要知道与它相连的所有节点以及它们之间的边的方向。
//我们使用一个 map 来表示这个图，其中 key 是节点编号，value 是一个数组，存储着与该节点相邻的节点编号以及它们之间的边的方向。
//如果从节点 i 到节点 j 的边是正向的，那么我们存储 j 到 i 的边的方向为负数，表示我们需要改变它们之间的方向。
//
//接下来，我们使用深度优先搜索算法来遍历图，统计需要改变方向的边的数量。我们从节点 0 开始遍历，访问每个相邻的节点。
//如果这个节点没有被访问过，我们就递归地遍历它，并计算需要改变方向的边的数量。
//如果当前节点与下一个节点之间的边需要改变方向，我们就将计数器加 1。
//
//最后，我们返回需要改变方向的边的数量作为结果。
//
//connections.length == n-1 说明图中没有环

func minReorder(n int, connections [][]int) int {
	// 构建图
	graph := make(map[int][]int)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], -conn[0]) // 表示改变方向
	}

	// DFS
	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(node int) int {
		count := 0
		visited[node] = true
		for _, next := range graph[node] {
			if !visited[abs(next)] {
				count += dfs(abs(next))
				if next > 0 {
					count++ // 需要改变方向
				}
			}
		}
		return count
	}

	// 统计需要改变方向的边的数量
	return dfs(0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

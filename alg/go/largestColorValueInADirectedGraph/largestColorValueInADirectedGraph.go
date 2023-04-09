// Source : https://leetcode.cn/problems/largest-color-value-in-a-directed-graph
// Date   : 2023-04-09

/**********************************************************************************
给你一个 有向图 ，它含有 n 个节点和 m 条边。节点编号从 0 到 n - 1 。
给你一个字符串 colors ，其中 colors[i] 是小写英文字母，表示图中第 i 个节点的 颜色 （下标从 0 开始）。同时给你一个二维数组 edges ，其中 edges[j] = [aj, bj] 表示从节点 aj 到节点 bj 有一条 有向边 。
图中一条有效 路径 是一个点序列 x1 -> x2 -> x3 -> ... -> xk ，对于所有 1 <= i < k ，从 xi 到 xi+1 在图中有一条有向边。路径的 颜色值 是路径中 出现次数最多 颜色的节点数目。
请你返回给定图中有效路径里面的 最大颜色值 。如果图中含有环，请返回 -1 。

示例 1：

输入：colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
输出：3
解释：路径 0 -> 2 -> 3 -> 4 含有 3 个颜色为 "a" 的节点（上图中的红色节点）。

示例 2：

输入：colors = "a", edges = [[0,0]]
输出：-1
解释：从 0 到 0 有一个环。


提示：

	n == colors.length
	m == edges.length
	1 <= n <= 105
	0 <= m <= 105
	colors 只含有小写英文字母。
	0 <= aj, bj < n
**********************************************************************************/

package largestColorValueInADirectedGraph

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	inDegree := make([]int, n)
	graph := make([][]int, n)

	// 建立邻接表并计算每个节点的入度
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	queue := make([]int, 0)   // 初始化BFS队列
	for i := range inDegree { // 将所有入度为0的节点入队
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	res := 0                   // 最终结果
	cnt := 0                   // 记录访问的节点数量
	visited := make([]bool, n) // 初始化节点访问状态
	dp := make([][26]int, n)   // dp[i]代表以i为路径结尾所出现的字符个数

	// BFS
	for len(queue) > 0 {
		from := queue[0] // 取出队首节点
		queue = queue[1:]

		cnt++                                      // 计数
		visited[from] = true                       // 标记节点已经被访问过
		dp[from][colors[from]-'a'] += 1            // 更新节点的字符出现次数
		res = max(res, dp[from][colors[from]-'a']) // 更新最终结果

		for _, to := range graph[from] { // 遍历所有的出边
			for c := range dp[from] { // dp状态转移方程
				dp[to][c] = max(dp[to][c], dp[from][c])
			}

			inDegree[to]--         // 更新后继节点的入度
			if inDegree[to] == 0 { // 如果入度为0，将其入队
				queue = append(queue, to)
			}
		}
	}

	// 如果有节点未被访问，说明有环
	if cnt != n {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

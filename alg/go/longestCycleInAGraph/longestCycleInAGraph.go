// Source : https://leetcode.cn/problems/longest-cycle-in-a-graph
// Date   : 2023-03-26

/**********************************************************************************
给你一个 n 个节点的 有向图 ，节点编号为 0 到 n - 1 ，其中每个节点 至多 有一条出边。
图用一个大小为 n 下标从 0 开始的数组 edges 表示，节点 i 到节点 edges[i] 之间有一条有向边。如果节点 i 没有出边，那么 edges[i] == -1 。
请你返回图中的 最长 环，如果没有任何环，请返回 -1 。
一个环指的是起点和终点是 同一个 节点的路径。

示例 1：


输入：edges = [3,3,4,2,3]
输出去：3
解释：图中的最长环是：2 -> 4 -> 3 -> 2 。
这个环的长度为 3 ，所以返回 3 。

示例 2：


输入：edges = [2,-1,3,1]
输出：-1
解释：图中没有任何环。


提示：

	n == edges.length
	2 <= n <= 105
	-1 <= edges[i] < n
	edges[i] != i
**********************************************************************************/

package longestCycleInAGraph

func longestCycle(edges []int) int {
	res := -1
	// 从每个节点开始遍历查找最长环
	for i := 0; i < len(edges); i++ {
		// 记录每个节点访问时的位置
		visited := make(map[int]int)
		res = max(dfs(i, 0, edges, visited), res)
	}
	return res
}

func dfs(node int, cur int, edges []int, visited map[int]int) int {
	// 表明无环
	if node == -1 {
		return -1
	}

	// 同一个节点第二次访问到，说明有环
	if pos, ok := visited[node]; ok {
		// 将本轮访问到的节点全部置为-1，避免后续重复计算（每个节点只有一个出边）
		for k := range visited {
			edges[k] = -1
		}
		return cur - pos
	}

	// 记录当前节点位置
	visited[node] = cur
	return dfs(edges[node], cur+1, edges, visited)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

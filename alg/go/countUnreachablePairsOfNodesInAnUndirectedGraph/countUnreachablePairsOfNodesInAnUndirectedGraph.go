// Source : https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph
// Date   : 2023-03-25

/**********************************************************************************
给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。
请你返回 无法互相到达 的不同 点对数目 。

示例 1：

输入：n = 3, edges = [[0,1],[0,2],[1,2]]
输出：0
解释：所有点都能互相到达，意味着没有点对无法互相到达，所以我们返回 0 。

示例 2：

输入：n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
输出：14
解释：总共有 14 个点对互相无法到达：
[[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6],[5,6]]
所以我们返回 14 。


提示：

	1 <= n <= 105
	0 <= edges.length <= 2 * 105
	edges[i].length == 2
	0 <= ai, bi < n
	ai != bi
	不会有重复边。
**********************************************************************************/

package countUnreachablePairsOfNodesInAnUndirectedGraph

type Graph struct {
	adj     [][]int
	visited []bool
}

func (g *Graph) dfs(u int) int {
	count := 1
	g.visited[u] = true
	for _, v := range g.adj[u] {
		if !g.visited[v] {
			count += g.dfs(v)
		}
	}
	return count
}

func countPairs(n int, edges [][]int) int64 {
	// 构建邻接表
	g := Graph{
		adj:     make([][]int, n),
		visited: make([]bool, n),
	}
	for _, e := range edges {
		u, v := e[0], e[1]
		g.adj[u] = append(g.adj[u], v)
		g.adj[v] = append(g.adj[v], u)
	}

	// 遍历整个图，获取不相交子图节点数
	var subs []int
	for i := 0; i < n; i++ {
		if !g.visited[i] {
			subs = append(subs, g.dfs(i))
		}
	}

	// 统计不可达的节点对数量
	count := 0
	for i := 0; i < len(subs); i++ {
		for j := i + 1; j < len(subs); j++ {
			count += subs[i] * subs[j]
		}
	}

	return int64(count)
}

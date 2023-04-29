// Source : https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths
// Date   : 2023-04-29

/**********************************************************************************
给你一个 n 个点组成的无向图边集 edgeList ，其中 edgeList[i] = [ui, vi, disi] 表示点 ui 和点 vi 之间有一条长度为 disi 的边。请注意，两个点之间可能有 超过一条边 。
给你一个查询数组queries ，其中 queries[j] = [pj, qj, limitj] ，你的任务是对于每个查询 queries[j] ，判断是否存在从 pj 到 qj 的路径，且这条路径上的每一条边都 严格小于 limitj 。
请你返回一个 布尔数组 answer ，其中 answer.length == queries.length ，当 queries[j] 的查询结果为 true 时， answer 第 j 个值为 true ，否则为 false 。

示例 1：

输入：n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
输出：[false,true]
解释： 2）两条边都小于 5 ，所以这个查询我们返回 true 。

示例 2：

输入：n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
输出：[true,false]
解释：上图为给定数据。


提示：

	2 <= n <= 105
	1 <= edgeList.length, queries.length <= 105
	edgeList[i].length == 3
	queries[j].length == 3
	0 <= ui, vi, pj, qj <= n - 1
	ui != vi
	pj != qj
	1 <= disi, limitj <= 109
	两个点之间可能有 多条 边。
**********************************************************************************/

package checkingExistenceOfEdgeLengthLimitedPaths

import "sort"

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	var union func(x, y int)
	union = func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			parent[x] = y
		}
	}

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	for i, query := range queries {
		queries[i] = append(query, i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})

	answer := make([]bool, len(queries))
	edgeIdx := 0
	for _, query := range queries {
		p, q, limit, idx := query[0], query[1], query[2], query[3]
		for edgeIdx < len(edgeList) && edgeList[edgeIdx][2] < limit {
			union(edgeList[edgeIdx][0], edgeList[edgeIdx][1])
			edgeIdx++
		}
		answer[idx] = find(p) == find(q)
	}

	return answer
}

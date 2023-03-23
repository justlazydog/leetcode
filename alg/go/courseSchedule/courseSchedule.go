// Source : https://leetcode.cn/problems/course-schedule
// Date   : 2023-03-23

/**********************************************************************************
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

	例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。

请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：

	1 <= numCourses <= 105
	0 <= prerequisites.length <= 5000
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	prerequisites[i] 中的所有课程对 互不相同
**********************************************************************************/

package courseSchedule

//算法使用拓扑排序的思想，先构建一个有向图，其中每个节点表示一个课程，每个先修关系表示一条有向边。
//然后，统计每个节点的入度，将入度为0的节点加入队列中，从队列中取出一个节点，并将该节点指向的节点的入度减1，如果入度减为0，则将该节点加入队列中。
//重复上述过程，直到队列为空。最后，如果存在入度不为0的节点，则说明存在环，课程表不合法，否则课程表合法。

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建入度表和邻接表
	inDegree := make([]int, numCourses)
	adjList := make([][]int, numCourses)
	for _, p := range prerequisites {
		inDegree[p[0]]++
		adjList[p[1]] = append(adjList[p[1]], p[0])
	}

	// 将入度为0的节点加入队列
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 依次删除入度为0的节点，并更新其邻接节点的入度
	count := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		count++
		for _, neighbor := range adjList[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 如果所有的节点都被删除，则表示可以完成课程
	return count == numCourses
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	// 构建邻接表
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
	}

	// 记录每个节点的访问状态
	visited := make([]int, numCourses)

	// 深度优先搜索遍历图
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !dfs(graph, visited, i) {
			return false
		}
	}

	return true
}

func dfs(graph [][]int, visited []int, i int) bool {
	// 当前节点已经被访问过了，说明存在环路
	if visited[i] == 1 {
		return false
	}

	// 标记当前节点已经被访问过
	visited[i] = 1

	// 遍历当前节点的所有邻居节点
	for _, neighbor := range graph[i] {
		if visited[neighbor] != 2 && !dfs(graph, visited, neighbor) {
			return false
		}
	}

	// 标记当前节点已经访问完毕
	visited[i] = 2

	return true
}

// Source : https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree
// Date   : 2023-04-19

/**********************************************************************************
给你一棵以 root 为根的二叉树，二叉树中的交错路径定义如下：

	选择二叉树中 任意 节点和一个方向（左或者右）。
	如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
	改变前进方向：左变右或者右变左。
	重复第二步和第三步，直到你在树中无法继续移动。

交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
请你返回给定树中最长 交错路径 的长度。

示例 1：

输入：root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
输出：3
解释：蓝色节点为树中最长交错路径（右 -> 左 -> 右）。

示例 2：

输入：root = [1,1,1,null,1,null,null,1,1,null,1]
输出：4
解释：蓝色节点为树中最长交错路径（左 -> 右 -> 左 -> 右）。

示例 3：
输入：root = [1]
输出：0


提示：

	每棵树最多有 50000 个节点。
	每个节点的值在 [1, 100] 之间。
**********************************************************************************/

package longestZigzagPathInABinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	maxLen := 0
	dfs(root, true, 0, &maxLen)
	dfs(root, false, 0, &maxLen)
	return maxLen
}

func dfs(node *TreeNode, isLeft bool, curLen int, maxLen *int) {
	if node == nil {
		return
	}

	if curLen > *maxLen {
		*maxLen = curLen
	}

	if isLeft {
		dfs(node.Right, !isLeft, curLen+1, maxLen)
		dfs(node.Right, isLeft, 0, maxLen)
	} else {
		dfs(node.Left, !isLeft, curLen+1, maxLen)
		dfs(node.Left, isLeft, 0, maxLen)
	}
}

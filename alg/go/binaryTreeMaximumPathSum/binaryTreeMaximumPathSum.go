// Source : https://leetcode.cn/problems/binary-tree-maximum-path-sum
// Date   : 2023-03-17

/**********************************************************************************
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

示例 1：

输入：root = [1,2,3]
输出：6
解释： 3 ，路径和为 2 + 1 + 3 = 6
示例 2：

输入：root = [-10,9,20,null,null,15,7]
输出：42
解释： 7 ，路径和为 15 + 20 + 7 = 42


提示：

	树中节点数目范围是 [1, 3 * 104]
	-1000 <= Node.val <= 1000
**********************************************************************************/

package binaryTreeMaximumPathSum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	maxPathSumHelper(root, &maxSum)
	return maxSum
}

func maxPathSumHelper(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	leftSum := maxPathSumHelper(node.Left, maxSum)
	rightSum := maxPathSumHelper(node.Right, maxSum)
	// 计算该节点的最大路径和
	nodeMaxSum := node.Val + max(0, leftSum) + max(0, rightSum)
	// 更新最大路径和
	*maxSum = max(*maxSum, nodeMaxSum)
	// 返回以该节点为根节点的最大贡献值（结点在路径中只能出现一次，所以贡献值只能是较大的子树提供）
	return node.Val + max(0, max(leftSum, rightSum))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

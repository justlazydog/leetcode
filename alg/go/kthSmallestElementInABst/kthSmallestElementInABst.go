// Source : https://leetcode.cn/problems/kth-smallest-element-in-a-bst
// Date   : 2023-03-25

/**********************************************************************************
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

示例 1：

输入：root = [3,1,4,null,2], k = 1
输出：1

示例 2：

输入：root = [5,3,6,2,4,null,null,1], k = 3
输出：3



提示：

	树中的节点数为 n 。
	1 <= k <= n <= 104
	0 <= Node.val <= 104


进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
**********************************************************************************/

package kthSmallestElementInABst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//由于二叉搜索树的中序遍历结果是一个有序序列，因此可以先对二叉搜索树进行中序遍历，
//然后在遍历过程中记录已经遍历的节点个数，当遍历到第 k 个节点时，该节点即为所求。

func kthSmallest(root *TreeNode, k int) int {
	var count int
	var res int
	// 定义中序遍历函数
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			res = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}

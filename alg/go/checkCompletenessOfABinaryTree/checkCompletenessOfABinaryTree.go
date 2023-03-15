// Source : https://leetcode.cn/problems/check-completeness-of-a-binary-tree
// Date   : 2023-03-15

/**********************************************************************************
给定一个二叉树的 root ，确定它是否是一个 完全二叉树 。
在一个 完全二叉树 中，除了最后一个关卡外，所有关卡都是完全被填满的，并且最后一个关卡中的所有节点都是尽可能靠左的。它可以包含 1 到 2h 节点之间的最后一级 h 。

示例 1：


输入：root = [1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。

示例 2：


输入：root = [1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。


提示：

	树的结点数在范围  [1, 100] 内。
	1 <= Node.val <= 1000
**********************************************************************************/

package checkCompletenessOfABinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	对一棵二叉树进行层序遍历时，我们先将根节点加入队列。
	接着从队列中取出第一个节点，
	如果该节点为空，则说明这棵树已经出现了空缺节点。
	如果该节点不为空且之前已经出现过空缺节点，则说明这棵树不是完全二叉树，我们直接返回 false。
	否则，将该节点的左右子节点加入队列。
	最后，如果所有节点都被访问过了，就说明这棵树是一棵完全二叉树，返回 true。
*/

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	var findEmpty bool
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			findEmpty = true
		} else {
			if findEmpty {
				return false
			}

			queue = append(queue, node.Left, node.Right)
		}
	}
	return true
}

// Source : https://leetcode.cn/problems/validate-binary-search-tree
// Date   : 2023-03-13

/**********************************************************************************
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：

	节点的左子树只包含 小于 当前节点的数。
	节点的右子树只包含 大于 当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。


示例 1：

输入：root = [2,1,3]
输出：true

示例 2：

输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。


提示：

	树中节点数目范围在[1, 104] 内
	-231 <= Node.val <= 231 - 1
**********************************************************************************/

package validateBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	一种简单的方法是中序遍历整个二叉树，然后判断遍历的结果是否是单调递增的。
	如果是单调递增的，那么该二叉树就是一棵二叉搜索树。
	这是因为中序遍历得到的结果就是树中节点值的升序排列。
*/

func isValidBST(root *TreeNode) bool {
	var lastVal int      // 保存中序遍历中上一个节点的值
	var firstNode = true // 当前节点是否是中序遍历的第一个节点
	return validate(root, &lastVal, &firstNode)
}

func validate(node *TreeNode, lastVal *int, firstNode *bool) bool {
	if node == nil {
		return true
	}

	if !validate(node.Left, lastVal, firstNode) {
		return false
	}

	if !*firstNode && node.Val <= *lastVal {
		return false
	}

	*firstNode = false
	*lastVal = node.Val

	if !validate(node.Right, lastVal, firstNode) {
		return false
	}

	return true
}

// Source : https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal
// Date   : 2023-03-16

/**********************************************************************************
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

示例 1:

输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:

输入：inorder = [-1], postorder = [-1]
输出：[-1]


提示:

	1 <= inorder.length <= 3000
	postorder.length == inorder.length
	-3000 <= inorder[i], postorder[i] <= 3000
	inorder 和 postorder 都由 不同 的值组成
	postorder 中每一个值都在 inorder 中
	inorder 保证是树的中序遍历
	postorder 保证是树的后序遍历
**********************************************************************************/

package constructBinaryTreeFromInorderAndPostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// similar problem construct-binary-tree-from-preorder-and-inorder-traversal

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var idx int
	for i, num := range inorder {
		if num == root.Val {
			idx = i
			break
		}
	}

	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}

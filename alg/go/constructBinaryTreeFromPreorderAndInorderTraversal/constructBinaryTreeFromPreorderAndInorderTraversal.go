// Source : https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// Date   : 2023-03-14

/**********************************************************************************
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:

输入: preorder = [-1], inorder = [-1]
输出: [-1]


提示:

	1 <= preorder.length <= 3000
	inorder.length == preorder.length
	-3000 <= preorder[i], inorder[i] <= 3000
	preorder 和 inorder 均 无重复 元素
	inorder 均出现在 preorder
	preorder 保证 为二叉树的前序遍历序列
	inorder 保证 为二叉树的中序遍历序列
**********************************************************************************/

package constructBinaryTreeFromPreorderAndInorderTraversal

/*
解题思路：

本题可以使用递归的方法解决。

前序遍历的第一个元素为根节点，在中序遍历中可以找到根节点的位置，这个位置将中序遍历数组分成了左右两个子数组。

对于前序遍历数组，我们需要根据根节点的位置确定左右子树的范围。然后递归地处理左右子树，直到数组为空。

在递归函数中，如果数组为空，返回 nil。
否则，创建一个根节点，并将前序遍历数组的第一个元素赋值给根节点。
然后找到根节点在中序遍历数组中的位置，并计算左右子树的范围。
在递归调用左右子树的过程中，需要更新前序遍历数组和中序遍历数组的指针和左右子树的范围。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	idx := 0
	for i, num := range inorder {
		if root.Val == num {
			idx = i
			break
		}
	}

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

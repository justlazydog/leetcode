// Source : https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree
// Date   : 2023-03-25

/**********************************************************************************
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

示例 1：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3 。

示例 2：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。

示例 3：

输入：root = [1,2], p = 1, q = 2
输出：1


提示：

	树中节点数目在范围 [2, 105] 内。
	-109 <= Node.val <= 109
	所有 Node.val 互不相同 。
	p != q
	p 和 q 均存在于给定的二叉树中。
**********************************************************************************/

package lowestCommonAncestorOfABinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//该解法采用递归的方式实现。首先判断根节点是否为空，或者是否等于给定的p或q节点，若是，则直接返回根节点。
//然后分别递归查找左子树和右子树，
//如果左子树和右子树均找到了p和q节点，则当前根节点即为最近公共祖先节点；
//如果只找到了左子树的p和q节点，则返回左子树的最近公共祖先节点；
//如果只找到了右子树的p和q节点，则返回右子树的最近公共祖先节点。

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

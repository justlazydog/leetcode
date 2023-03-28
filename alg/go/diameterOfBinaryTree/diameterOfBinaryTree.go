// Source : https://leetcode.cn/problems/diameter-of-binary-tree
// Date   : 2023-03-28

/**********************************************************************************
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5

返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
**********************************************************************************/

package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDistance int
	dfs(root, &maxDistance) // 调用 dfs 函数计算二叉树的直径长度
	return maxDistance
}

// 这条路径可能穿过也可能不穿过根结点。
func dfs(node *TreeNode, maxDistance *int) int {
	if node == nil {
		return 0 // 如果节点为空，返回深度为 0
	}

	leftDepth := dfs(node.Left, maxDistance)   // 递归计算左子树的深度
	rightDepth := dfs(node.Right, maxDistance) // 递归计算右子树的深度

	if leftDepth+rightDepth > *maxDistance { // 如果左子树和右子树深度之和大于当前的直径长度
		*maxDistance = leftDepth + rightDepth // 更新直径长度
	}

	return max(leftDepth, rightDepth) + 1 // 返回当前节点为根节点的子树的深度
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

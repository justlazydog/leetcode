// Source : https://leetcode.cn/problems/sum-root-to-leaf-numbers
// Date   : 2023-03-14

/**********************************************************************************
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。

每条从根节点到叶节点的路径都代表一个数字：

	例如，从根节点到叶节点的路径  3 表示数字 123 。

计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。

示例 1：

输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 2 代表数字 12
从根到叶子节点路径 3 代表数字 13
因此，数字总和 = 12 + 13 = 25
示例 2：

输入：root = [4,9,0,5,1]
输出：1026
解释：
从根到叶子节点路径 5 代表数字 495
从根到叶子节点路径 1 代表数字 491
从根到叶子节点路径 0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026


提示：

	树中节点的数目在范围 [1, 1000] 内
	0 <= Node.val <= 9
	树的深度不超过 10

**********************************************************************************/

//package sumRootToLeafNumbers

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var res int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, pathSum int, res *int) {
	if root == nil {
		return
	}

	pathSum = pathSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		*res += pathSum
		return
	}

	dfs(root.Left, pathSum, res)
	dfs(root.Right, pathSum, res)
}

//func dfs(root *TreeNode, tmp string, res *[]string) {
//	if root == nil {
//		return
//	}
//
//	tmp += strconv.Itoa(root.Val)
//
//	if root.Left == nil && root.Right == nil {
//		*res = append(*res, tmp)
//		tmp = tmp[:len(tmp)-1]
//	}
//
//	dfs(root.Left, tmp, res)
//	dfs(root.Right, tmp, res)
//	if len(tmp) > 0 {
//		tmp = tmp[:len(tmp)-1]
//	}
//}

//func main() {
//	root := &TreeNode{
//		Val: 4,
//	}
//	node1 := &TreeNode{
//		Val: 9,
//	}
//	node2 := &TreeNode{
//		Val: 5,
//	}
//	node3 := &TreeNode{
//		Val: 1,
//	}
//	node4 := &TreeNode{
//		Val: 0,
//	}
//	root.Left = node1
//	root.Right = node4
//	node1.Left = node2
//	node1.Right = node3
//	var res []string
//	dfs(root, "", &res)
//	fmt.Println(res)
//}

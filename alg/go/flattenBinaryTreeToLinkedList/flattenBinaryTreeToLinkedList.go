// Source : https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
// Date   : 2023-03-14

/**********************************************************************************
给你二叉树的根结点 root ，请你将它展开为一个单链表：

	展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
	展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：

输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：

输入：root = []
输出：[]

示例 3：

输入：root = [0]
输出：[0]


提示：

	树中结点数在范围 [0, 2000] 内
	-100 <= Node.val <= 100


进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
**********************************************************************************/

package flattenBinaryTreeToLinkedList

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	具体地，将二叉树展开为一个单链表的步骤如下：

	将左子树展开为链表，并返回链表的末尾节点，即左子树的最后一个节点。
	将右子树展开为链表，并返回链表的末尾节点，即右子树的最后一个节点。
	连接左右子树：如果左子树不为空，则将左子树的末尾节点的右子节点指向右子树的根节点，并将根节点的左子节点指向空；如果左子树为空，则不需要做任何操作。
	返回链表的末尾节点：如果右子树不为空，则链表的末尾节点为右子树的末尾节点；如果右子树为空且左子树不为空，则链表的末尾节点为左子树的末尾节点；
	如果左右子树都为空，则链表的末尾节点为根节点。
*/

func flatten(root *TreeNode) {
	flattenTree(root)
}

func flattenTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	// 左子树展开为链表
	leftTail := flattenTree(node.Left)
	// 右子树展开为链表
	rightTail := flattenTree(node.Right)

	// 连接左右子树
	if leftTail != nil {
		leftTail.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}

	// 返回链表的末尾
	if rightTail != nil {
		return rightTail
	} else if leftTail != nil {
		return leftTail
	} else {
		return node
	}
}

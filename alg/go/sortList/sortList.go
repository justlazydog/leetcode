// Source : https://leetcode.cn/problems/sort-list
// Date   : 2023-03-21

/**********************************************************************************
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。



示例 1：

输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：

输入：head = []
输出：[]


提示：

	链表中节点的数目在范围 [0, 5 * 104] 内
	-105 <= Node.val <= 105


进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
**********************************************************************************/

package sortList

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表的中间节点
	var slow, fast = head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	// 对左右两个子链表分别进行排序
	left := sortList(head)
	right := sortList(mid)

	// 合并左右两个已经排好序的子链表
	var dummy, cur = &ListNode{Val: 0}, &ListNode{Val: 0}
	dummy = cur
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return dummy.Next
}

// Source : https://leetcode.cn/problems/palindrome-linked-list
// Date   : 2023-03-25

/**********************************************************************************
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：

输入：head = [1,2,2,1]
输出：true

示例 2：

输入：head = [1,2]
输出：false


提示：

	链表中节点数目在范围[1, 105] 内
	0 <= Node.val <= 9


进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
**********************************************************************************/

package palindromeLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 如果链表为空或只有一个节点，直接返回 true
	if head == nil || head.Next == nil {
		return true
	}
	// 使用快慢指针找到链表的中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转链表的后半部分
	var prev *ListNode
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 将反转后的后半部分与原链表的前半部分进行比较
	p1, p2 := head, prev
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

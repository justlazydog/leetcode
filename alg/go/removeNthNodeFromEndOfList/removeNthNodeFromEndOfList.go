// Source : https://leetcode.cn/problems/remove-nth-node-from-end-of-list
// Date   : 2023-03-07

/**********************************************************************************
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：

输入：head = [1], n = 1
输出：[]

示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

	链表中结点的数目为 sz
	1 <= sz <= 30
	0 <= Node.val <= 100
	1 <= n <= sz


进阶：你能尝试使用一趟扫描实现吗？
**********************************************************************************/

package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

//下面是一份使用双指针的解法：
//
//1. 定义两个指针p和q，分别指向链表的头节点。
//
//2. 将p指针向前移动N个节点，然后将q指针和p指针一起向前移动，直到p指针到达链表的末尾。
//
//3. 此时q指针指向要删除的节点的前一个节点，将其next指针指向要删除的节点的后一个节点即可。
//
//4. 如果要删除的节点是链表的头节点，需要特殊处理。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head

	for i := 0; i < n; i++ {
		p = p.Next
	}

	if p == nil {
		return head.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	q.Next = q.Next.Next
	return head
}

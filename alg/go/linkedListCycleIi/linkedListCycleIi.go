// Source : https://leetcode.cn/problems/linked-list-cycle-ii
// Date   : 2023-03-17

/**********************************************************************************
给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
不允许修改 链表。



示例 1：


输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：


输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：


输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。


提示：

	链表中节点的数目范围在范围 [0, 104] 内
	-105 <= Node.val <= 105
	pos 的值为 -1 或者链表中的一个有效索引


进阶：你是否可以使用 O(1) 空间解决此题？
**********************************************************************************/

package linkedListCycleIi

type ListNode struct {
	Val  int
	Next *ListNode
}

//解题思路：
//
//使用双指针，首先判断链表是否存在环，如果存在环，找到环的位置。找到环后，使用双指针方法找到环的入口。
//
//步骤如下：
//
//设置快慢指针，分别指向头结点。
//
//当快指针不为 null 时，慢指针移动一步，快指针移动两步。
//
//如果快指针到达链表尾部，则说明链表不存在环，直接返回 null。
//
//如果快指针和慢指针相遇，则说明链表存在环，此时快指针重新指向头结点，并以每次移动一步的速度移动，慢指针保持不变。
//
//当快指针和慢指针相遇时，相遇的节点就是环的起点。

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

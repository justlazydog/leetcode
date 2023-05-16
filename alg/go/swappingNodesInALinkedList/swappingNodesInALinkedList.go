// Source : https://leetcode.cn/problems/swapping-nodes-in-a-linked-list
// Date   : 2023-05-15

/**********************************************************************************
给你链表的头节点 head 和一个整数 k 。
交换 链表正数第 k 个节点和倒数第 k 个节点的值后，返回链表的头节点（链表 从 1 开始索引）。

示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[1,4,3,2,5]

示例 2：

输入：head = [7,9,6,6,7,8,3,0,9,5], k = 5
输出：[7,9,6,6,8,7,3,0,9,5]

示例 3：

输入：head = [1], k = 1
输出：[1]

示例 4：

输入：head = [1,2], k = 1
输出：[2,1]

示例 5：

输入：head = [1,2,3], k = 2
输出：[1,2,3]


提示：

	链表中节点的数目是 n
	1 <= k <= n <= 105
	0 <= Node.val <= 100
**********************************************************************************/

//package swappingNodesInALinkedList

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 创建链表
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}

	// 构建链表关系
	head.Next = node1
	node1.Next = node2
	node2.Next = node3

	// 打印链表
	printList(swapNodes(head, 1))
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func swapNodes(head *ListNode, k int) *ListNode {
	//dummy := &ListNode{Next: head}
	slow, fast := head, head
	count := 1
	for count < k && fast.Next != nil {
		count++
		fast = fast.Next
	}
	node1 := fast

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	node2 := slow
	node1.Val, node2.Val = node2.Val, node1.Val

	return head
}

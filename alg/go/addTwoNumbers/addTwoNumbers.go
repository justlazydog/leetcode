// Source : https://leetcode.cn/problems/add-two-numbers
// Date   : 2022-12-05

/**********************************************************************************
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例 1：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：

	每个链表中的节点数在范围 [1, 100] 内
	0
	题目数据保证列表表示的数字不含前导零
**********************************************************************************/

package main

import (
	"math/rand"
	"time"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, res := 0, new(ListNode)
	for node := res; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return res.Next
}

// generateLinkedList generates a linked list of random length between minLen and maxLen
// with random integer values between minValue and maxValue.
func generateLinkedList(minLen, maxLen, minValue, maxValue int) *ListNode {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(maxLen-minLen+1) + minLen
	head := &ListNode{}

	for i := 0; i < n; i++ {
		val := rand.Intn(maxValue-minValue+1) + minValue
		node := &ListNode{Val: val, Next: head.Next}
		head.Next = node
	}

	return head.Next
}

// recursion
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersHelper(l1, l2, 0)
}

func addTwoNumbersHelper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	v1, v2 := 0, 0
	if l1 != nil {
		v1 = l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		v2 = l2.Val
		l2 = l2.Next
	}

	sum := v1 + v2 + carry
	carry = sum / 10
	node := &ListNode{Val: sum % 10}
	node.Next = addTwoNumbersHelper(l1, l2, carry)
	if node.Next == nil && carry != 0 {
		node.Next = &ListNode{Val: carry}
	}
	return node
}

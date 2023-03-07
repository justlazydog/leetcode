// Source : https://leetcode.cn/problems/merge-k-sorted-lists
// Date   : 2023-03-07

/**********************************************************************************
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]

示例 3：
输入：lists = [[]]
输出：[]


提示：

	k == lists.length
	0 <= k <= 10^4
	0 <= lists[i].length <= 500
	-10^4 <= lists[i][j] <= 10^4
	lists[i] 按 升序 排列
	lists[i].length 的总和不超过 10^4
**********************************************************************************/

package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PQ, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	if len(pq) == 0 {
		return nil
	}
	heap.Init(&pq)

	head := &ListNode{}
	dummyHead := head

	for len(pq) > 0 {
		node := heap.Pop(&pq).(*ListNode)
		head.Next = node
		head = head.Next

		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return dummyHead.Next
}

type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].Val < pq[b].Val
}

func (pq *PQ) Push(nodeInterface interface{}) {
	node := nodeInterface.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	lastNode := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return lastNode
}

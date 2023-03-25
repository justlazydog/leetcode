// Source : https://leetcode.cn/problems/kth-largest-element-in-an-array
// Date   : 2023-03-25

/**********************************************************************************
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：

	1 <= k <= nums.length <= 105
	-104 <= nums[i] <= 104
**********************************************************************************/

package kthLargestElementInAnArray

import "container/heap"

type queue []int

func (q queue) Len() int {
	return len(q)
}

func (q queue) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *queue) Pop() interface{} {
	n := len(*q)
	item := (*q)[n-1]
	*q = (*q)[:n-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	h := &queue{}
	heap.Init(h)
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else if num > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}
	return (*h)[0]
}

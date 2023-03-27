// Source : https://leetcode.cn/problems/top-k-frequent-elements
// Date   : 2023-03-27

/**********************************************************************************
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:

输入: nums = [1], k = 1
输出: [1]

提示：

	1 <= nums.length <= 105
	k 的取值范围是 [1, 数组中不相同的元素的个数]
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的


进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
**********************************************************************************/

package topKFrequentElements

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	// 统计每个元素出现的次数
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 将哈希表中的元素转化为元素数组
	n := len(freqMap)
	elems := make([][2]int, n)
	i := 0
	for num, freq := range freqMap {
		elems[i] = [2]int{num, freq}
		i++
	}

	// 对元素数组进行排序，按照元素出现的次数从大到小排列
	sort.Slice(elems, func(i, j int) bool {
		return elems[i][1] > elems[j][1]
	})

	// 返回排序后的前 k 个元素
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = elems[i][0]
	}
	return res
}

//我们也可以使用堆来解决这个问题，我们可以将哈希表中的元素逐个加入堆中，并维护一个大小为 k 的小根堆，堆中元素是元素出现的次数以及元素本身。
//如果当前堆的大小超过了 k，就弹出堆顶元素。最后返回堆中剩余的 k 个元素。
//
//具体步骤如下：
//
//使用哈希表统计每个元素出现的次数；
//将哈希表中的元素逐个加入堆中，维护一个大小为 k 的小根堆，堆中元素是元素出现的次数以及元素本身；
//如果当前堆的大小超过了 k，就弹出堆顶元素；
//返回堆中

func topKFrequent2(nums []int, k int) []int {
	// 统计每个元素出现的次数
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 将哈希表中的元素逐个加入堆中，维护一个大小为 k 的小根堆
	h := &IntHeap{}
	heap.Init(h)
	for num, freq := range freqMap {
		heap.Push(h, [2]int{freq, num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 返回堆中剩余的 k 个元素
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).([2]int)[1]
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Source : https://leetcode.cn/problems/find-median-from-data-stream
// Date   : 2023-03-26

/**********************************************************************************
中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。

	例如 arr = [2,3,4] 的中位数是 3 。
	例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。

实现 MedianFinder 类:


	MedianFinder() 初始化 MedianFinder 对象。


	void addNum(int num) 将数据流中的整数 num 添加到数据结构中。


	double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。


示例 1：

输入
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
输出
[null, null, null, 1.5, null, 2.0]
解释
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
提示:

	-105 <= num <= 105
	在调用 findMedian 之前，数据结构中至少有一个元素
	最多 5 * 104 次调用 addNum 和 findMedian
**********************************************************************************/

package findMedianFromDataStream

//这里我们使用两个堆来实现，一个小根堆用来存储所有大于等于中位数的数，一个大根堆用来存储所有小于中位数的数。
//当两个堆的元素个数相差超过1时，我们需要将多余的元素移动到另一个堆中，保证两个堆的元素个数相差不超过1。
//
//在每次插入元素时，我们需要判断该元素应该插入哪个堆中。
//如果该元素小于等于大根堆的堆顶元素，则插入大根堆中；否则，插入小根堆中。
//然后，我们需要判断两个堆的元素个数是否平衡，如果不平衡，则需要将多余的元素移动到另一个堆中。
//
//最后，我们可以根据两个堆的元素个数来计算中位数。
//如果两个堆的元素个数不相等，则中位数为元素个数多的那个堆的堆顶元素；否则，中位数为两个堆的堆顶元素之和除以2。

type MedianFinder struct {
	minHeap *Heap // 堆顶元素最小
	maxHeap *Heap // 堆顶元素最大
}

// Constructor initializes the MedianFinder data structure.
func Constructor() MedianFinder {
	minHeap := &Heap{make([]int, 0), func(i, j int) bool { return i < j }}
	maxHeap := &Heap{make([]int, 0), func(i, j int) bool { return i > j }}
	return MedianFinder{minHeap, maxHeap}
}

// AddNum adds a number into the MedianFinder data structure.
func (mf *MedianFinder) AddNum(num int) {
	if mf.maxHeap.Len() == 0 || num <= mf.maxHeap.Peek() {
		mf.maxHeap.Push(num)
	} else {
		mf.minHeap.Push(num)
	}
	if mf.maxHeap.Len()-mf.minHeap.Len() > 1 {
		mf.minHeap.Push(mf.maxHeap.Pop())
	} else if mf.minHeap.Len()-mf.maxHeap.Len() > 1 {
		mf.maxHeap.Push(mf.minHeap.Pop())
	}
}

// FindMedian finds the median of the MedianFinder data structure.
func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64(mf.maxHeap.Peek())
	} else if mf.minHeap.Len() > mf.maxHeap.Len() {
		return float64(mf.minHeap.Peek())
	} else {
		return float64(mf.maxHeap.Peek()+mf.minHeap.Peek()) / 2.0
	}
}

// Heap is a simple implementation of a binary heap.
type Heap struct {
	arr  []int
	less func(i, j int) bool
}

// Len returns the number of elements in the heap.
func (h *Heap) Len() int {
	return len(h.arr)
}

// Push adds a new element to the heap.
func (h *Heap) Push(x int) {
	h.arr = append(h.arr, x)
	h.siftUp(h.Len() - 1)
}

// Pop removes and returns the top element of the heap.
func (h *Heap) Pop() int {
	n := h.Len() - 1
	h.arr[0], h.arr[n] = h.arr[n], h.arr[0]
	h.siftDown(0, n)
	x := h.arr[n]
	h.arr = h.arr[:n]
	return x
}

// Peek returns the top element of the heap without removing it.
func (h *Heap) Peek() int {
	return h.arr[0]
}

// siftUp moves the element at index i up the heap until it is in the correct position.
func (h *Heap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.less(h.arr[i], h.arr[p]) {
			h.arr[i], h.arr[p] = h.arr[p], h.arr[i]
			i = p
		} else {
			break
		}
	}
}

// siftDown moves the element at index i down the heap until it is in the correct position.
func (h *Heap) siftDown(i, n int) {
	for i*2+1 < n {
		j := i*2 + 1
		if j+1 < n && h.less(h.arr[j+1], h.arr[j]) {
			j++
		}
		if h.less(h.arr[j], h.arr[i]) {
			h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
			i = j
		} else {
			break
		}
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

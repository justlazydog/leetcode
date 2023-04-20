// Source : https://leetcode.cn/problems/sort-an-array
// Date   : 2023-04-20

/**********************************************************************************
给你一个整数数组 nums，请你将该数组升序排列。



示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]

示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]


提示：

	1 <= nums.length <= 5 * 104
	-5 * 104 <= nums[i] <= 5 * 104
**********************************************************************************/

package sortAnArray

func sortArray(nums []int) []int {
	// 构建最大堆
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	// 逐个将堆顶元素移到末尾，重新调整堆
	for i := n - 1; i >= 0; i-- {
		// 将堆顶元素与末尾元素交换位置
		nums[0], nums[i] = nums[i], nums[0]

		// 调整堆
		heapify(nums, i, 0)
	}
	return nums
}

func heapify(arr []int, n, i int) {
	largest := i // 初始化最大值为根节点
	l := 2*i + 1 // 左子节点的下标
	r := 2*i + 2 // 右子节点的下标

	// 找出左右子节点中的最大值
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// 如果最大值不是根节点，则交换根节点和最大值，并继续调整子树
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

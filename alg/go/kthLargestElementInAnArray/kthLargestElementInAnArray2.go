package kthLargestElementInAnArray

func findKthLargest2(nums []int, k int) int {
	// 快速选择算法
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	// 选取最后一个元素作为枢纽元素
	pivot := nums[right]
	i := left
	// 将大于主元素的元素放在左边，小于等于的元素放在右边
	for j := left; j < right; j++ {
		if nums[j] >= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将主元素放在正确的位置
	nums[i], nums[right] = nums[right], nums[i]

	// 如果枢纽元素恰好是第 k 大的元素，直接返回
	if i-left+1 == k {
		return nums[i]
	}
	// 如果枢纽元素所在的位置比 k 大，说明目标元素在左边
	if i-left+1 > k {
		return quickSelect(nums, left, i-1, k)
	}
	// 如果枢纽元素所在的位置比 k 小，说明目标元素在右边
	return quickSelect(nums, i+1, right, k-(i-left+1))
}

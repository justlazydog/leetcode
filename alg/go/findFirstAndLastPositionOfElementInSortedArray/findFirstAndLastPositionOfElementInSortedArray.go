// Source : https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
// Date   : 2023-03-09

/**********************************************************************************
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：

	0 <= nums.length <= 105
	-109 <= nums[i] <= 109
	nums 是一个非递减数组
	-109 <= target <= 109
**********************************************************************************/

package findFirstAndLastPositionOfElementInSortedArray

func searchRange(nums []int, target int) []int {
	i := binarySearch(nums, target, 0, len(nums)-1)
	left, right := i, i
	for left > 0 && nums[left-1] == target {
		left--
	}

	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{left, right}
}

func binarySearch(nums []int, target int, low, high int) int {
	if len(nums) == 0 {
		return -1
	}

	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return binarySearch(nums, target, low, mid-1)
	} else {
		return binarySearch(nums, target, mid+1, high)
	}
}

// Source : https://leetcode.cn/problems/first-missing-positive
// Date   : 2023-03-09

/**********************************************************************************
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：

输入：nums = [1,2,0]
输出：3

示例 2：

输入：nums = [3,4,-1,1]
输出：2

示例 3：

输入：nums = [7,8,9,11,12]
输出：1


提示：

	1 <= nums.length <= 5 * 105
	-231 <= nums[i] <= 231 - 1
**********************************************************************************/

package firstMissingPositive

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 标记所有小于等于 0 和大于数组长度的数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}
	// 将所有正整数对应的数组元素标记为负数表示该数出现过
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	// 找到第一个未被标记的位置
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Source : https://leetcode.cn/problems/maximum-subarray
// Date   : 2023-03-11

/**********************************************************************************
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：

输入：nums = [1]
输出：1

示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

	1 <= nums.length <= 105
	-104 <= nums[i] <= 104


进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
**********************************************************************************/

package maximumSubarray

func maxSubArray(nums []int) int {
	curSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

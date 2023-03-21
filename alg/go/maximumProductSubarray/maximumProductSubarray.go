// Source : https://leetcode.cn/problems/maximum-product-subarray
// Date   : 2023-03-21

/**********************************************************************************
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
子数组 是数组的连续子序列。

示例 1:

输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:

输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

提示:

	1 <= nums.length <= 2 * 104
	-10 <= nums[i] <= 10
	nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
**********************************************************************************/

package maximumProductSubarray

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 初始化结果变量
	maxProduct, minProduct, ans := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		// 如果nums[i]是负数，那么maxProduct和minProduct交换
		if nums[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		// 更新maxProduct和minProduct
		maxProduct = max(maxProduct*nums[i], nums[i])
		minProduct = min(minProduct*nums[i], nums[i])

		// 更新ans
		ans = max(ans, maxProduct)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

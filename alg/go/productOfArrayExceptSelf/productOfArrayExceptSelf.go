// Source : https://leetcode.cn/problems/product-of-array-except-self
// Date   : 2023-03-25

/**********************************************************************************
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请不要使用除法，且在 O(n) 时间复杂度内完成此题。

示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]

示例 2:

输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]


提示：

	2 <= nums.length <= 105
	-30 <= nums[i] <= 30
	保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内


进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
**********************************************************************************/

package productOfArrayExceptSelf

//在这个解法中，我们首先创建一个长度为 n 的数组 res，用于存放结果。
//然后，我们使用 res[i] 存储 nums[i] 左侧所有元素的乘积，从而避免了每次进行重复计算。
//接着，我们再计算 nums[i] 右侧所有元素的乘积，将其与 res[i] 相乘，即可得到 nums 中除了 nums[i] 以外的所有元素的乘积。

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// res[i] 表示 nums[i] 左侧所有元素的乘积
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// right 表示 nums[i] 右侧所有元素的乘积
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

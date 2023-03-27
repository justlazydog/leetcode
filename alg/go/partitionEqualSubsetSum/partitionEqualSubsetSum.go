// Source : https://leetcode.cn/problems/partition-equal-subset-sum
// Date   : 2023-03-27

/**********************************************************************************
给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

示例 1：

输入：nums = [1,5,11,5]
输出：true
解释：数组可以分割成 [1, 5, 5] 和 [11] 。
示例 2：

输入：nums = [1,2,3,5]
输出：false
解释：数组不能分割成两个元素和相等的子集。


提示：

	1 <= nums.length <= 200
	1 <= nums[i] <= 100
**********************************************************************************/

package partitionEqualSubsetSum

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 { // 如果数组元素的和不能被 2 整除，则不可能划分成两个元素和相等的子集
		return false
	}
	target := sum / 2 // 计算出数组元素和的一半

	dp := make([]bool, target+1) // 创建长度为 target+1 的布尔型数组，dp[i] 表示是否能够组成和为 i 的子集
	dp[0] = true                 // 前 0 个数能够组成和为 0 的子集

	for i := 1; i <= len(nums); i++ { // 遍历前 i 个数
		for j := target; j >= nums[i-1]; j-- { // 从 target 到 nums[i-1] 遍历
			dp[j] = dp[j] || dp[j-nums[i-1]] // 状态转移方程
		}
	}

	return dp[target] // 返回 dp[target]
}

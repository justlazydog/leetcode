// Source : https://leetcode.cn/problems/subarray-sum-equals-k
// Date   : 2023-03-28

/**********************************************************************************
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。

示例 1：

输入：nums = [1,1,1], k = 2
输出：2

示例 2：

输入：nums = [1,2,3], k = 3
输出：2


提示：

	1 <= nums.length <= 2 * 104
	-1000 <= nums[i] <= 1000
	-107 <= k <= 107
**********************************************************************************/

package subarraySumEqualsK

func subarraySum(nums []int, k int) int {
	count, preSum := 0, 0
	hash := make(map[int]int)
	hash[0] = 1 // 初始化哈希表，因为前缀和为 0 的情况有一种

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if v, ok := hash[preSum-k]; ok {
			count += v // 如果存在前缀和等于 preSum - k 的情况，则将其数量加入计数器
		}
		hash[preSum]++ // 将当前前缀和加入哈希表中
	}

	return count
}

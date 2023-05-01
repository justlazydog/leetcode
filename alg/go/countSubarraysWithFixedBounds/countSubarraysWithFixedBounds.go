// Source : https://leetcode.cn/problems/count-subarrays-with-fixed-bounds
// Date   : 2023-05-01

/**********************************************************************************
给你一个整数数组 nums 和两个整数 minK 以及 maxK 。
nums 的定界子数组是满足下述条件的一个子数组：

	子数组中的 最小值 等于 minK 。
	子数组中的 最大值 等于 maxK 。

返回定界子数组的数目。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [1,3,5,2,7,5], minK = 1, maxK = 5
输出：2
解释：定界子数组是 [1,3,5] 和 [1,3,5,2] 。

示例 2：
输入：nums = [1,1,1,1], minK = 1, maxK = 1
输出：10
解释：nums 的每个子数组都是一个定界子数组。共有 10 个子数组。

提示：

	2 <= nums.length <= 105
	1 <= nums[i], minK, maxK <= 106
**********************************************************************************/

package countSubarraysWithFixedBounds

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64
	var lastMin, lastMax, left = -1, -1, 0

	for i, num := range nums {
		if num < minK || num > maxK {
			left = i + 1
			lastMin = -1
			lastMax = -1
			continue
		}

		if num == minK {
			lastMin = i
		}

		if num == maxK {
			lastMax = i
		}

		if lastMin != -1 && lastMax != -1 {
			res += int64(min(lastMin, lastMax) - left + 1)
		}
	}

	return res
}

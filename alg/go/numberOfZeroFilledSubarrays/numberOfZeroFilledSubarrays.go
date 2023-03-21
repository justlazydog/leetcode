// Source : https://leetcode.cn/problems/number-of-zero-filled-subarrays
// Date   : 2023-03-21

/**********************************************************************************
给你一个整数数组 nums ，返回全部为 0 的 子数组 数目。
子数组 是一个数组中一段连续非空元素组成的序列。

示例 1：
输入：nums = [1,3,0,0,2,0,0,4]
输出：6
解释：
子数组 [0] 出现了 4 次。
子数组 [0,0] 出现了 2 次。
不存在长度大于 2 的全 0 子数组，所以我们返回 6 。
示例 2：
输入：nums = [0,0,0,2,0,0]
输出：9
解释：
子数组 [0] 出现了 5 次。
子数组 [0,0] 出现了 3 次。
子数组 [0,0,0] 出现了 1 次。
不存在长度大于 3 的全 0 子数组，所以我们返回 9 。

示例 3：
输入：nums = [2,10,2019]
输出：0
解释：没有全 0 子数组，所以我们返回 0 。


提示：

	1 <= nums.length <= 105
	-109 <= nums[i] <= 109
**********************************************************************************/

package numberOfZeroFilledSubarrays

func zeroFilledSubarray(nums []int) int64 {
	// easy deal nums[len-1] == 0
	nums = append(nums, 1)
	var ans, zeros int64 = 0, 0
	for _, n := range nums {
		if n == 0 {
			zeros++
		} else if zeros > 0 {
			ans += zeros * (zeros + 1) / 2
			zeros = 0
		}
	}

	return ans
}

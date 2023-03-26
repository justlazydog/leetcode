// Source : https://leetcode.cn/problems/find-the-duplicate-number
// Date   : 2023-03-26

/**********************************************************************************
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：

输入：nums = [1,3,4,2,2]
输出：2

示例 2：

输入：nums = [3,1,3,4,2]
输出：3


提示：

	1 <= n <= 105
	nums.length == n + 1
	1 <= nums[i] <= n
	nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次


进阶：

	如何证明 nums 中至少存在一个重复的数字?
	你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？
**********************************************************************************/

package findTheDuplicateNumber

//我们可以利用抽屉原理（鸽笼原理）来优化这个方法。
//抽屉原理是指如果有 n 个物品放进 m 个抽屉中，且 n>m，则至少有一个抽屉中放了两个或两个以上的物品。
//
//对于这道题，数组中每个元素的值都在 1 到 n 之间，那么我们可以将数组下标看作鸽笼，数组中的元素看作物品。
//由于数组中至少存在一个重复的整数，那么一定有一个鸽笼中至少放了两个或两个以上的物品，也就是说，数组下标存在重复。

func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 1, n-1
	for left < right {
		mid := left + (right-left)/2
		count := 0

		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

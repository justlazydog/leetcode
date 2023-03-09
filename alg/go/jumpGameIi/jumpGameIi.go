// Source : https://leetcode.cn/problems/jump-game-ii
// Date   : 2023-03-09

/**********************************************************************************
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

	0 <= j <= nums[i]
	i + j < n

返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:

输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:

输入: nums = [2,3,0,1,4]
输出: 2


提示:

	1 <= nums.length <= 104
	0 <= nums[i] <= 1000
	题目保证可以到达 nums[n-1]
**********************************************************************************/

package jumpGameIi

/*
	在函数 jump 中，
	maxPos 表示当前能够跳跃的最大步数，
	end 表示当前能够跳跃的最大步数所能够到达的最远位置，
	step 表示跳跃的步数。

	在循环中，
	更新 maxPos 的值为 max(maxPos, i+nums[i])，其中 i+nums[i] 表示当前位置所能够到达的最远位置。
	如果 maxPos 能够直接跳到最后位置，直接 step + 1 退出循环，
	如果 i 等于 end，说明需要再跳一步才能继续更新当前能够跳跃的最大步数，此时跳跃的步数加 1，同时将 end 的值更新为 maxPos。

	最后返回跳跃的步数 step。
*/

func jump(nums []int) int {
	n := len(nums)
	maxPos, end, step := 0, 0, 0
	for i := 0; i < n-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if maxPos >= n-1 {
			step++
			break
		}

		if i == end {
			end = maxPos
			step++
		}

	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

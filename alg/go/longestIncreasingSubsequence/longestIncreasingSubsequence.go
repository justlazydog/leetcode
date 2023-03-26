// Source : https://leetcode.cn/problems/longest-increasing-subsequence
// Date   : 2023-03-26

/**********************************************************************************
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4

示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1


提示：

	1 <= nums.length <= 2500
	-104 <= nums[i] <= 104


进阶：

	你能将算法的时间复杂度降低到 O(n log(n)) 吗?
**********************************************************************************/

package longestIncreasingSubsequence

//我们定义一个长度为 n 的数组 dp，其中 dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度。
//初始化时，我们将所有的 dp[i] 都设置为 1，因为每个数都可以看作是一个长度为 1 的递增子序列。
//
//然后，我们从前往后遍历数组 nums，对于每个元素 nums[i]，我们遍历它之前的所有元素 nums[j]，
//如果 nums[j] < nums[i]，则表示 nums[i] 可以加入 nums[j] 对应的递增子序列中，从而得到一个更长的递增子序列，更新 dp[i] 的值。
//具体而言，我们可以将 dp[i] 的值更新为 dp[j]+1，表示将 nums[i] 加入到以 nums[j] 结尾的递增子序列中。
//这里我们还需要维护一个 maxLen 变量，记录 dp 数组中的最大值，即最长的递增子序列的长度。

func lengthOfLIS(nums []int) int {
	// dp[i] 表示以 nums[i] 结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//我们定义一个长度为 n 的数组 tails，其中 tails[i] 表示长度为 i+1 的递增子序列的最后一个元素的最小值。
//在遍历数组 nums 的过程中，我们维护一个变量 len 表示当前的递增子序列长度。
//
//对于每个元素 nums[i]，我们需要在 tails 数组中查找一个位置 j，使得 tails[j-1] < nums[i] <= tails[j]， 然后更新 tails[j] 为 nums[i]。
//如果 nums[i] 大于 tails 数组中的最大值，则将 nums[i] 直接添加到 tails 数组的末尾， 并将 len 加 1。
//否则，我们使用二分查找在 tails 数组中查找合适的位置，从而使得我们能够更新 tails 数组中的元素，保证其仍然是一个递增的子序列。
//具体而言，我们可以在 tails 数组中查找第一个大于等于 nums[i] 的位置，将其更新为 nums[i]。
//https://leetcode.com/problems/longest-increasing-subsequence/solutions/74824/javapython-binary-search-onlogn-time-with-explanation/

func lengthOfLIS2(nums []int) int {
	tails := make([]int, len(nums))
	tails[0] = nums[0]
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[maxLen-1] {
			tails[maxLen] = nums[i]
			maxLen++
		} else {
			l, r := 0, maxLen-1
			for l < r {
				mid := l + (r-l)/2
				if tails[mid] < nums[i] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			tails[l] = nums[i]
		}
	}

	return maxLen
}

// Source : https://leetcode.cn/problems/coin-change
// Date   : 2023-03-27

/**********************************************************************************
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0


提示：

	1 <= coins.length <= 12
	1 <= coins[i] <= 231 - 1
	0 <= amount <= 104
**********************************************************************************/

package coinChange

import (
	"math"
	"sort"
)

//解法一：动态规划
//首先我们定义一个数组 dp，其中 dp[i] 表示组成总金额为 i 的最小硬币数目。
//对于任意一个总金额 i，我们遍历硬币面额 coins[j]，如果 coins[j] <= i，则表示当前这个面额的硬币可以选择，我们可以从 dp[i-coins[j]] 转移过来，得到 dp[i] 的值。
//在所有的 dp[i-coins[j]] 中选择最小的一个，再加上当前这个硬币，即可得到 dp[i] 的值。

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

//贪心算法 + 剪枝
//我们可以先将硬币面额 coins 按照从大到小的顺序排序，然后从大到小依次枚举硬币面额 coins[j]，并将其加入到组成总金额 amount 的方案中。
//我们定义一个变量 cnt，表示已经选取的硬币数目，同时记录剩余的总金额为 rest。
//对于当前的硬币面额 coins[j]，我们计算能够选取的最大硬币数目 k，即 k = rest / coins[j]。
//然后，我们从 k 开始递减枚举，将 coins[j]加入到组成总金额的方案中，同时更新 cnt 和 rest 的值。
//如果在更新的过程中发现当前已选取的硬币数目 cnt 已经大于了当前的最优解 ans，那么就可以剪枝，直接结束搜索.
//如果 rest 的值已经为 0，说明当前方案是可行的，比较 cnt 和当前最优解 ans 的大小，更新 ans 的值。
//最终，如果 ans 的值仍然为正无穷大，说明无法组成总金额，返回 -1，否则返回 ans 的值。

func coinChange2(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	ans := math.MaxInt32

	var dfs func(cnt, rest, idx int)
	dfs = func(cnt, rest, idx int) {
		if rest == 0 {
			ans = min(ans, cnt)
			return
		}
		if idx == len(coins) {
			return
		}
		k := rest / coins[idx]
		for i := k; i >= 0 && cnt+i < ans; i-- {
			dfs(cnt+i, rest-i*coins[idx], idx+1)
		}
	}

	dfs(0, amount, 0)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

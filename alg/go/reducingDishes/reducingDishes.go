// Source : https://leetcode.cn/problems/reducing-dishes
// Date   : 2023-03-29

/**********************************************************************************
一个厨师收集了他 n 道菜的满意程度 satisfaction ，这个厨师做出每道菜的时间都是 1 单位时间。
一道菜的 「喜爱时间」系数定义为烹饪这道菜以及之前每道菜所花费的时间乘以这道菜的满意程度，也就是 time[i]*satisfaction[i] 。
请你返回做完所有菜 「喜爱时间」总和的最大值为多少。
你可以按 任意 顺序安排做菜的顺序，你也可以选择放弃做某些菜来获得更大的总和。

示例 1：

输入：satisfaction = [-1,-8,0,5,-9]
输出：14
解释：去掉第二道和最后一道菜，最大的喜爱时间系数和为 (-1*1 + 0*2 + 5*3 = 14) 。每道菜都需要花费 1 单位时间完成。
示例 2：

输入：satisfaction = [4,3,2]
输出：20
解释：按照原来顺序相反的时间做菜 (2*1 + 3*2 + 4*3 = 20)

示例 3：

输入：satisfaction = [-1,-4,-5]
输出：0
解释：大家都不喜欢这些菜，所以不做任何菜可以获得最大的喜爱时间系数。


提示：

	n == satisfaction.length
	1 <= n <= 500
	-1000 <= satisfaction[i] <= 1000
**********************************************************************************/

package reducingDishes

import "sort"

func maxSatisfaction(satisfaction []int) int {
	var res int
	sort.Ints(satisfaction) // 将满意度数组按照从小到大排序

	// 如果第一个元素就是非负数，则全部选中，顺序为1, 2, 3, ..., n
	if satisfaction[0] >= 0 {
		for i, s := range satisfaction {
			res += (i + 1) * s
		}
		return res
	}

	n := len(satisfaction)
	// 如果最后一个元素就是非正数，则全部不选
	if satisfaction[n-1] <= 0 {
		return 0
	}

	// 动态规划求解
	dp := make([]int, n+1)
	dp[1] = satisfaction[n-1] // 只选最后一个元素的情况
	for i := 2; i <= n; i++ {
		if satisfaction[n-i] < 0 { // 如果当前元素为负数
			tmp := 0
			for j := i; j > 0; j-- { // 找到最大的正数子序列
				tmp += (i - j + 1) * satisfaction[n-j]
			}
			dp[i] = max(dp[i-1], tmp)
		} else { // 如果当前元素为非负数
			dp[i] = satisfaction[n-i] + dp[i-1] // 考虑选当前元素的情况
			for j := i; j > 1; j-- {            // 继续选前面的正数元素
				dp[i] += satisfaction[n-(j-1)]
			}
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

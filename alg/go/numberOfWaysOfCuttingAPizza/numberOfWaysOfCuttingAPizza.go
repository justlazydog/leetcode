// Source : https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza
// Date   : 2023-03-31

/**********************************************************************************
给你一个 rows x cols 大小的矩形披萨和一个整数 k ，矩形包含两种字符： &#39;A&#39; （表示苹果）和 &#39;.&#39; （表示空白格子）。你需要切披萨 k-1 次，得到 k 块披萨并送给别人。
切披萨的每一刀，先要选择是向垂直还是水平方向切，再在矩形的边界上选一个切的位置，将披萨一分为二。如果垂直地切披萨，那么需要把左边的部分送给一个人，如果水平地切，那么需要把上面的部分送给一个人。在切完最后一刀后，需要把剩下来的一块送给最后一个人。
请你返回确保每一块披萨包含 至少 一个苹果的切披萨方案数。由于答案可能是个很大的数字，请你返回它对 10^9 + 7 取余的结果。

示例 1：

输入：pizza = ["A..","AAA","..."], k = 3
输出：3
解释：上图展示了三种切披萨的方案。注意每一块披萨都至少包含一个苹果。

示例 2：
输入：pizza = ["A..","AA.","..."], k = 3
输出：1

示例 3：
输入：pizza = ["A..","A..","..."], k = 1
输出：1


提示：

	1 <= rows, cols <= 50
	rows == pizza.length
	cols == pizza[i].length
	1 <= k <= 10
	pizza 只包含字符 &#39;A&#39; 和 &#39;.&#39; 。
**********************************************************************************/

package numberOfWaysOfCuttingAPizza

func ways(pizza []string, k int) int {
	m, n := len(pizza), len(pizza[0])
	// dp[i][j][s] 表示从矩阵的 (i,j) 位置开始切 s 刀，使得每个子矩阵中至少包含一个苹果的切法总数。
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// appleSum 变量是一个二维数组，表示从矩阵的 (i,j) 位置开始到矩阵的右下角的子矩阵中苹果的总数。
	appleSum := make([][]int, m+1)
	for i := range appleSum {
		appleSum[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			appleSum[i][j] = appleSum[i+1][j] + appleSum[i][j+1] - appleSum[i+1][j+1]
			if pizza[i][j] == 'A' {
				appleSum[i][j]++
			}
		}
	}

	// Initialize the base cases
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if appleSum[i][j] > 0 {
				dp[i][j][0] = 1
			}
		}
	}

	// Compute the number of ways to cut the pizza
	for s := 1; s < k; s++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for x := i + 1; x < m; x++ {
					// 如果从 (i,j) 到 (x-1,j) 的子矩阵中有苹果
					if appleSum[i][j]-appleSum[x][j] > 0 {
						dp[i][j][s] += dp[x][j][s-1]
						dp[i][j][s] %= 1e9 + 7
					}
				}
				for y := j + 1; y < n; y++ {
					// 如果从 (i,j) 到 (i,y-1) 的子矩阵中有苹果
					if appleSum[i][j]-appleSum[i][y] > 0 {
						dp[i][j][s] += dp[i][y][s-1]
						dp[i][j][s] %= 1e9 + 7
					}
				}
			}
		}
	}

	return dp[0][0][k-1]
}

// Source : https://leetcode.cn/problems/edit-distance
// Date   : 2023-03-13

/**********************************************************************************
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：

	插入一个字符
	删除一个字符
	替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

	0 <= word1.length, word2.length <= 500
	word1 和 word2 由小写英文字母组成
**********************************************************************************/

package editDistance

/*
	这是一道经典的动态规划问题，可以使用动态规划算法解决。

	动态规划算法的基本思路是，用一个二维数组dp[i][j]表示将第一个字符串的前i个字符转换为第二个字符串的前j个字符所需的最少编辑操作次数。根据不同的情况，可以进行不同的转移：

	如果第一个字符串的第i个字符和第二个字符串的第j个字符相同，则dp[i][j]等于dp[i-1][j-1]，不需要进行任何编辑操作；

	如果第一个字符串的第i个字符和第二个字符串的第j个字符不同，则有以下三种操作方式：
	将第一个字符串的第i个字符替换为第二个字符串的第j个字符，需要执行一次替换操作，所以dp[i][j]等于dp[i-1][j-1]+1；
	在第一个字符串的第i个字符后面插入第二个字符串的第j个字符，需要执行一次插入操作，所以dp[i][j]等于dp[i][j-1]+1；
	删除第一个字符串的第i个字符，需要执行一次删除操作，所以dp[i][j]等于dp[i-1][j]+1。

	最终返回dp[m][n]，其中m和n分别是两个字符串的长度。
*/

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			}
		}
	}

	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

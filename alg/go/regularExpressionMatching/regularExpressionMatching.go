// Source : https://leetcode.cn/problems/regular-expression-matching
// Date   : 2023-04-10

/**********************************************************************************
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素

所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

示例 1：

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。

示例 2:

输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3：

输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。


提示：

	1 <= s.length <= 20
	1 <= p.length <= 20
	s 只包含从 a-z 的小写字母。
	p 只包含从 a-z 的小写字母，以及字符 . 和 *。
	保证每次出现字符 * 时，前面都匹配到有效的字符
**********************************************************************************/

package regularExpressionMatching

// 判断字符串 s 和模式串 p 是否匹配
func isMatch(s string, p string) bool {
	// 计算字符串 s 和模式串 p 的长度
	m, n := len(s), len(p)

	// 初始化 dp 数组
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// 空字符串和空模式串匹配
	dp[0][0] = true

	// 计算 dp 数组中的其他位置
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 如果模式串的第 j-1 个字符是 *，则可以有两种选择：
				// 选择不匹配 p[j-2]*，则 dp[i][j] = dp[i][j-2]
				dp[i][j] = dp[i][j-2]

				// 选择匹配 p[j-2]*，则需要判断 s[i-1] 和 p[j-2] 是否匹配
				// 如果匹配，则 dp[i][j] = dp[i][j] || dp[i-1][j]
				if match(s, p, i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				// 如果模式串的第 j-1 个字符不是 *，则只需要判断 s[i-1] 和 p[j-1] 是否相等
				// 如果相等，则 dp[i][j] = dp[i-1][j-1]
				if match(s, p, i, j) {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	// 返回 dp[m][n]，表示字符串 s 和模式串 p 是否匹配
	return dp[m][n]
}

// 判断字符串 s 和模式串 p 的第 i 和第 j 个字符是否匹配
func match(s string, p string, i int, j int) bool {
	if i == 0 {
		// 如果 s 已经匹配完了，但是 p 还没有匹配完，返回 false
		return false
	}
	if p[j-1] == '.' {
		// 如果模式串的第 j-1 个字符是 .，则可以匹配任意字符
		return true
	}
	// 否则，只有当 s[i-1] 和 p[j-1] 相等时，才匹配成功
	return s[i-1] == p[j-1]
}

// Source : https://leetcode.cn/problems/number-of-ways-to-form-a-target-string-given-a-dictionary
// Date   : 2023-04-16

/**********************************************************************************
给你一个字符串列表 words 和一个目标字符串 target 。words 中所有字符串都 长度相同  。
你的目标是使用给定的 words 字符串列表按照下述规则构造 target ：

	从左到右依次构造 target 的每一个字符。
	为了得到 target 第 i 个字符（下标从 0 开始），当 target[i] = words[j][k] 时，你可以使用 words 列表中第 j 个字符串的第 k 个字符。
	一旦你使用了 words 中第 j 个字符串的第 k 个字符，你不能再使用 words 字符串列表中任意单词的第 x 个字符（x <= k）。也就是说，所有单词下标小于等于 k 的字符都不能再被使用。
	请你重复此过程直到得到目标字符串 target 。

请注意， 在构造目标字符串的过程中，你可以按照上述规定使用 words 列表中 同一个字符串 的 多个字符 。
请你返回使用 words 构造 target 的方案数。由于答案可能会很大，请对 109 + 7 取余 后返回。
（译者注：此题目求的是有多少个不同的 k 序列，详情请见示例。）

示例 1：

输入：words = ["acca","bbbb","caca"], target = "aba"
输出：6
解释： 下标为 0 ("acca")，下标为 1 ("bbbb")，下标为 3 ("caca 下标为 0 ("acca")，下标为 2 ("bbbb")，下标为 3 ("caca 下标为 0 ("acca")，下标为 1 ("bbbb")，下标为 3 ("acca 下标为 0 ("acca")，下标为 2 ("bbbb")，下标为 3 ("acca 下标为 1 ("caca")，下标为 2 ("bbbb")，下标为 3 ("acca 下标为 1 ("caca")，下标为 2 ("bbbb")，下标为 3 ("caca")

示例 2：

输入：words = ["abba","baab"], target = "bab"
输出：4
解释： 下标为 0 ("baab")，下标为 1 ("baab")，下标为 2 ("abb 下标为 0 ("baab")，下标为 1 ("baab")，下标为 3 ("baab 下标为 0 ("baab")，下标为 2 ("baab")，下标为 3 ("baab 下标为 1 ("abba")，下标为 2 ("baab")，下标为 3 ("baab")

示例 3：

输入：words = ["abcd"], target = "abcd"
输出：1

示例 4：

输入：words = ["abab","baba","abba","baab"], target = "abba"
输出：16


提示：

	1 <= words.length <= 1000
	1 <= words[i].length <= 1000
	words 中所有单词长度相同。
	1 <= target.length <= 1000
	words[i] 和 target 都仅包含小写英文字母。
**********************************************************************************/

package numberOfWaysToFormATargetStringGivenADictionary

func numWays(words []string, target string) int {
	mod := int(1e9 + 7)
	m := len(words[0])
	targetLength := len(target)

	// 计算每个位置上字母出现的频率
	freq := make([][26]int, m)
	for _, word := range words {
		for i, ch := range word {
			freq[i][ch-'a']++
		}
	}

	// 初始化dp数组
	// dp[i][j] 表示从第i个字符开始，使用频率数组中第j列及之后的列来组成目标字符串的方案数
	dp := make([][]int, targetLength)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// 动态规划计算方案数
	for i := targetLength - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == targetLength-1 {
				if j == m-1 {
					dp[i][j] = freq[j][target[i]-'a']
				} else {
					dp[i][j] = (dp[i][j+1] + freq[j][target[i]-'a']) % mod
				}
			} else {
				if j != m-1 {
					dp[i][j] = (dp[i][j+1] + dp[i+1][j+1]*freq[j][target[i]-'a']) % mod
				}
			}
		}
	}

	return dp[0][0]
}

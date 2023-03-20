// Source : https://leetcode.cn/problems/word-break
// Date   : 2023-03-20

/**********************************************************************************
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。

示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


提示：

	1 <= s.length <= 300
	1 <= wordDict.length <= 1000
	1 <= wordDict[i].length <= 20
	s 和 wordDict[i] 仅有小写英文字母组成
	wordDict 中的所有字符串 互不相同
**********************************************************************************/

package wordBreak

//定义一个长度为 n+1 的布尔数组 dp，其中 dp[i] 表示字符串的前 i 个字符是否可以被空格拆分成若干个在 wordDict 中出现的单词。
//例如，当 s = "leetcode"，wordDict = ["leet", "code"] 时，dp[4] = true，因为字符串的前 4 个字符 "leet" 出现在 wordDict 中。
//
//我们可以通过动态规划来求解 dp 数组。初始时，dp[0] 为 true，表示空字符串可以被任何单词拆分。
//对于 i 从 1 到 n，我们枚举字符串的每个前缀，如果前缀是一个在 wordDict 中出现的单词，并且剩余部分也可以被拆分，那么字符串前 i 个字符就可以被拆分。

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

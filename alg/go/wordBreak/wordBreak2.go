package wordBreak

//我们也可以使用深度优先搜索来解决这个问题。具体来说，我们从字符串的开头开始搜索，枚举当前单词的结束位置，
//如果当前单词存在于 wordDict 中，那么我们将搜索指针移动到当前单词的下一个位置，继续搜索剩余的字符串。
//如果搜索指针移动到字符串的末尾，那么就说明字符串可以被拆分成若干个在 wordDict 中出现的单词。
//
//为了避免重复计算，我们可以使用一个数组 memo 来记录搜索指针从某个位置开始搜索的结果。
//具体来说，数组 memo 中的第 i 个元素表示从字符串的第 i 个位置开始搜索的结果，
//如果值为 true，那么从这个位置开始的字符串可以被拆分成若干个在 wordDict 中出现的单词，如果值为 false，则表示不行。

func wordBreak2(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	memo := make([]int8, n)

	var dfs func(int) bool
	dfs = func(start int) bool {
		if start == n {
			return true
		}

		if memo[start] != 0 {
			return memo[start] == 1
		}

		for end := start + 1; end <= n; end++ {
			if wordSet[s[start:end]] && dfs(end) {
				memo[start] = 1
				return true
			}
		}

		memo[start] = -1
		return false
	}

	return dfs(0)
}

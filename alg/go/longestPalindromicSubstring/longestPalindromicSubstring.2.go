// Source : https://leetcode.cn/problems/longest-palindromic-substring
// Date   : 2022-12-05

/**********************************************************************************
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：

输入：s = "cbbd"
输出："bb"


提示：

	1 <= s.length <= 1000
	s 仅由数字和英文字母组成
**********************************************************************************/

package longestPalindrome

// longestPalindromeDP is solution by dynamic programming
func longestPalindromeDP(s string) string {
	var start, end int

	for i := 0; i < len(s); i++ {
		left, right := i, i

		// 跳过重复字符（重复字符必是回文串）
		for right < len(s)-1 && s[left] == s[right+1] {
			right++
		}

		// 向两边扩撒匹配，若匹配即是回文串
		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		// 是否是最长
		if right-left > end-start {
			start = left
			end = right
		}
	}

	return s[start : end+1]
}

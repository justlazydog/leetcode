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

// brute force

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}
	return true
}

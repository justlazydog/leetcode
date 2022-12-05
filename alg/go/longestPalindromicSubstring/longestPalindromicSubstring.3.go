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

// longestPalindromeEAC is solution by expand around center
func longestPalindromeEAC(s string) string {
	var res, str string

	// 长度为n的字符，中心数为2n-1
	centerNum := len(s)*2 - 1

	for i := 1; i <= centerNum; i++ {
		if i%2 == 1 {
			str = getPalindromeOddCenter(i, s)
		} else {
			str = getPalindromeEvenCenter(i, s)
		}

		if len(str) > len(res) {
			res = str
		}
	}

	return res
}

// getPalindromeOddCenter get palindrome from odd center
// a b b c
// ^ ^ ^ ^
// 1 3 5 7
func getPalindromeOddCenter(center int, s string) string {
	left := (center - 1) / 2
	right := left

	for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
		left--
		right++
	}

	return s[left : right+1]
}

// getPalindromeOddCenter get palindrome from even center
// a b b c
//
//	^ ^ ^
//	2 4 6
func getPalindromeEvenCenter(center int, s string) string {
	left := center/2 - 1
	right := left + 1

	if s[left] != s[right] {
		return s[left:right]
	}

	for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
		left--
		right++
	}

	return s[left : right+1]
}

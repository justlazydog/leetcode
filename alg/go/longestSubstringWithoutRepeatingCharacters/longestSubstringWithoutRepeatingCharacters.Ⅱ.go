// Source : https://leetcode.cn/problems/longest-substring-without-repeating-characters
// Date   : 2022-12-05

/**********************************************************************************
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


提示：

	0 <= s.length <= 5 * 104
	s 由英文字母、数字、符号和空格组成
**********************************************************************************/

package lengthOfLongestSubstring

// lengthOfLongestSubstring2 is the same idea as lengthOfLongestSubstring but faster
func lengthOfLongestSubstring2(s string) int {
	res := 0
	m := make(map[rune]int)

	left := 0
	for right, c := range s {
		if _, ok := m[c]; ok {
			if m[c] > left {
				left = m[c]
			}
		}

		l := right - left + 1
		if l > res {
			res = l
		}

		m[c] = right + 1
	}

	return res
}

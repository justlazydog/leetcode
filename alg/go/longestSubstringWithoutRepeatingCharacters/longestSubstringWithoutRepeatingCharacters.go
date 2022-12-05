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

import (
	"strings"
)

/*

	s:
	abcdabcbb

	str:
	abcd
	 bcda
	  cdab
	   dabc
		  cb
			b

	longest = max(str)

*/

func lengthOfLongestSubstring(s string) int {
	var (
		str     string
		longest int
	)

	for i := 0; i < len(s); i++ {
		index := strings.IndexByte(str, s[i])
		if index >= 0 {
			if len(s)-index <= longest {
				break
			}
			str = str[index+1:]
		}
		str += string(s[i])

		if len(str) > longest {
			longest = len(str)
		}
	}

	return longest
}

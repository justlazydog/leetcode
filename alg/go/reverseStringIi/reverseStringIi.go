// Source : https://leetcode.cn/problems/reverse-string-ii
// Date   : 2023-04-29

/**********************************************************************************
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

	如果剩余字符少于 k 个，则将剩余字符全部反转。
	如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。


示例 1：

输入：s = "abcdefg", k = 2
输出："bacdfeg"

示例 2：

输入：s = "abcd", k = 2
输出："bacd"


提示：

	1 <= s.length <= 104
	s 仅由小写英文组成
	1 <= k <= 104
**********************************************************************************/

package reverseStringIi

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i += 2 * k {
		left, right := i, i+k-1
		if right >= len(bytes) {
			right = len(bytes) - 1
		}
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}

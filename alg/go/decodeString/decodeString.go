// Source : https://leetcode.cn/problems/decode-string
// Date   : 2023-03-27

/**********************************************************************************
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"


提示：

	1 <= s.length <= 30
	s 由小写英文字母、数字和方括号 '[]' 组成
	s 保证是一个 有效 的输入。
	s 中所有整数的取值范围为 [1, 300]
**********************************************************************************/

package decodeString

import "strings"

func decodeString(s string) string {
	var numStack []int    // 存储数字的栈
	var strStack []string // 存储字符串的栈
	num := 0              // 当前解码的数字
	cur := ""             // 当前正在解码的字符串

	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0') // 读取数字，并计算出总数
		} else if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			cur += string(c) // 将字母添加到当前正在解码的字符串中
		} else if c == '[' {
			numStack = append(numStack, num) // 将当前解码的数字压入数字栈中
			strStack = append(strStack, cur) // 将当前正在解码的字符串压入字符串栈中
			num, cur = 0, ""                 // 重置 num 和 cur
		} else if c == ']' {
			k := numStack[len(numStack)-1]        // 取出数字栈顶元素 k
			numStack = numStack[:len(numStack)-1] // 将数字栈顶元素弹出
			prev := strStack[len(strStack)-1]     // 取出字符串栈顶元素 prev
			strStack = strStack[:len(strStack)-1] // 将字符串栈顶元素弹出
			cur = prev + strings.Repeat(cur, k)   // 将 cur 重复 k 次，并与 prev 拼接
		}
	}

	return cur // 栈中只剩下一个字符串，它就是解码后的结果
}

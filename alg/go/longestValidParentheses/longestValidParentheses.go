// Source : https://leetcode.cn/problems/longest-valid-parentheses
// Date   : 2023-03-08

/**********************************************************************************
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。


示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：

输入：s = ""
输出：0


提示：

	0 <= s.length <= 3 * 104
	s[i] 为 '(' 或 ')'

**********************************************************************************/

package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	stack = append(stack, -1) // 用-1初始化栈，方便计算有效括号序列长度

	maxLen := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i) // 左括号入栈
		} else {
			stack = stack[:len(stack)-1] // 右括号弹出栈顶元素
			if len(stack) == 0 {
				stack = append(stack, i) // 如果栈为空，将右括号入栈，更新起始点
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1]) // 计算有效括号序列长度
			}
		}
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("()"))
}

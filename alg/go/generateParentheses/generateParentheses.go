// Source : https://leetcode.cn/problems/generate-parentheses
// Date   : 2023-03-07

/**********************************************************************************
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：

输入：n = 1
输出：["()"]


提示：

	1 <= n <= 8
**********************************************************************************/

package generateParenthesis

func generateParenthesis(n int) []string {
	var res []string
	dfs(n, n, "", &res)
	return res
}

func dfs(left, right int, s string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}

	if left > 0 {
		dfs(left-1, right, s+"(", res)
	}

	if right > 0 && right > left {
		dfs(left, right-1, s+")", res)
	}
}

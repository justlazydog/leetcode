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
	res := make([]string, 0)
	// left,right mean open and close brackets num
	var dfs func(left, right int, cur string)
	dfs = func(left, right int, cur string) {
		if left == 0 && right == 0 {
			res = append(res, cur)
			return
		}
		if left > 0 {
			dfs(left-1, right, cur+"(")
		}
		if right > 0 && right > left {
			dfs(left, right-1, cur+")")
		}
	}
	dfs(n, n, "")
	return res
}

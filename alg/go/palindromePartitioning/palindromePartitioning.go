// Source : https://leetcode.cn/problems/palindrome-partitioning
// Date   : 2023-03-19

/**********************************************************************************
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

示例 2：

输入：s = "a"
输出：[["a"]]


提示：

	1 <= s.length <= 16
	s 仅由小写英文字母组成
**********************************************************************************/

package palindromePartitioning

//为了解决这个问题，我们可以使用回溯法。具体来说，我们定义一个辅助函数 backtrack(start)，其中 start 表示当前考虑到的位置。该函数的执行流程为：
//
//如果 start 等于字符串 s 的长度，则找到了一种分割方案，将其加入答案数组中。
//否则，枚举子串的结束位置 i，如果子串 s[start..i] 是回文串，则将其加入路径列表 path 中，并在此基础上继续回溯，回溯完成后将子串 s[start..i] 从 path 中删除，然后尝试下一个结束位置。

func partition(s string) [][]string {
	var res [][]string
	var path []string
	backtrack(0, s, path, &res)
	return res
}

func backtrack(start int, s string, path []string, res *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			backtrack(i+1, s, append(path, s[start:i+1]), res)
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

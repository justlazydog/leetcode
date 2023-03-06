// Source : https://leetcode.cn/problems/longest-common-prefix
// Date   : 2023-03-06

/**********************************************************************************
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] 仅由小写英文字母组成
**********************************************************************************/

package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	// find min length in strs
	min := len(strs[0])
	for _, str := range strs {
		if min > len(str) {
			min = len(str)
		}
	}

	// compare prefix
	var i int
	for i = 0; i < min; i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if c != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:i]
}

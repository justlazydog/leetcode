// Source : https://leetcode.cn/problems/scramble-string
// Date   : 2023-03-30

/**********************************************************************************
使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
	如果字符串的长度为 1 ，算法停止
	 1 ，执行下述步骤：

		在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
		随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
		在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。



给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。

示例 1：

输入：s1 = "great", s2 = "rgeat"
输出：true
解释： "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
算法终止，结果字符串和 s2 相同，都是 "rgeat"
这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true

示例 2：

输入：s1 = "abcde", s2 = "caebd"
输出：false

示例 3：

输入：s1 = "a", s2 = "a"
输出：true


提示：

	s1.length == s2.length
	1 <= s1.length <= 30
	s1 和 s2 由小写英文字母组成
**********************************************************************************/

package scrambleString

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) { // 字符串长度不同，不能通过扰乱得到
		return false
	}
	if s1 == s2 { // 字符串相等，可以通过扰乱得到
		return true
	}

	// 检查s1和s2是否由相同的字母组成
	count := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}

	n := len(s1)
	for i := 1; i < n; i++ {
		// 检查s1[:i]和s2[:i]是否可以通过扰乱得到，s1[i:]和s2[i:]是否可以通过扰乱得到
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		// 检查s1[:i]和s2[n-i:]是否可以通过扰乱得到，s1[i:]和s2[:n-i]是否可以通过扰乱得到
		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[:n-i]) {
			return true
		}
	}
	return false
}

// Source : https://leetcode.cn/problems/find-all-anagrams-in-a-string
// Date   : 2023-03-28

/**********************************************************************************
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


提示:

	1 <= s.length, p.length <= 3 * 104
	s 和 p 仅包含小写字母
**********************************************************************************/

package findAllAnagramsInAString

func findAnagrams(s string, p string) []int {
	var res []int
	n, m := len(s), len(p)
	if n < m {
		return res
	}

	freqP, freqS := make([]int, 26), make([]int, 26)
	for i := 0; i < m; i++ {
		freqP[p[i]-'a']++
	}

	left, right := 0, 0
	for right < n {
		freqS[s[right]-'a']++
		if right-left+1 > m {
			freqS[s[left]-'a']--
			left++
		}
		if right-left+1 == m && isAnagram(freqS, freqP) {
			res = append(res, left)
			freqS[s[left]-'a']--
			left++
		}
		right++
	}

	return res
}

func isAnagram(freqS []int, freqP []int) bool {
	for i := 0; i < 26; i++ {
		if freqS[i] != freqP[i] {
			return false
		}
	}
	return true
}

// Source : https://leetcode.cn/problems/permutation-in-string
// Date   : 2023-03-28

/**********************************************************************************
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false


提示：

	1 <= s1.length, s2.length <= 104
	s1 和 s2 仅包含小写字母
**********************************************************************************/

package permutationInString

func checkInclusion(s1 string, s2 string) bool {
	return findAnagrams(s2, s1)
}

func findAnagrams(s2 string, s1 string) bool {
	n, m := len(s2), len(s1)
	if n < m {
		return false
	}

	freqP, freqS := make([]int, 26), make([]int, 26)
	for i := 0; i < m; i++ {
		freqP[s1[i]-'a']++
	}

	left, right := 0, 0
	for right < n {
		freqS[s2[right]-'a']++
		if right-left+1 > m {
			freqS[s2[left]-'a']--
			left++
		}
		if right-left+1 == m && isAnagram(freqS, freqP) {
			return true
		}
		right++
	}

	return false
}

func isAnagram(freqS []int, freqP []int) bool {
	for i := 0; i < 26; i++ {
		if freqS[i] != freqP[i] {
			return false
		}
	}
	return true
}

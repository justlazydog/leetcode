// Source : https://leetcode.cn/problems/find-common-characters
// Date   : 2023-04-28

/**********************************************************************************
给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序 返回答案。

示例 1：

输入：words = ["bella","label","roller"]
输出：["e","l","l"]

示例 2：

输入：words = ["cool","lock","cook"]
输出：["c","o"]


提示：

	1 <= words.length <= 100
	1 <= words[i].length <= 100
	words[i] 由小写英文字母组成
**********************************************************************************/

package findCommonCharacters

func commonChars(words []string) []string {
	var res []string
	if len(words) == 0 {
		return res
	}

	cnt := [26]int{}
	for _, ch := range words[0] {
		cnt[ch-'a']++
	}

	for i := 1; i < len(words); i++ {
		cur := [26]int{}
		for _, ch := range words[i] {
			cur[ch-'a']++
		}
		for j := 0; j < 26; j++ {
			cnt[j] = min(cnt[j], cur[j])
		}
	}

	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			res = append(res, string('a'+byte(i)))
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

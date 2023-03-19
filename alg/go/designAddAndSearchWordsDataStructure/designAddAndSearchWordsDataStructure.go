// Source : https://leetcode.cn/problems/design-add-and-search-words-data-structure
// Date   : 2023-03-19

/**********************************************************************************
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
实现词典类 WordDictionary ：

	WordDictionary() 初始化词典对象
	void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
	bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。


示例：

输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]
解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True
wordDictionary.search("b.."); // 返回 True


提示：

	1 <= word.length <= 25
	addWord 中的 word 由小写英文字母组成
	search 中的 word 由 '.' 或小写英文字母组成
	最多调用 104 次 addWord 和 search
**********************************************************************************/

package designAddAndSearchWordsDataStructure

type WordDictionary struct {
	isEnd bool
	next  map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{next: make(map[byte]*WordDictionary)}
}

func (d *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		d.isEnd = true
		return
	}

	firstChar := word[0]
	if d.next[firstChar] == nil {
		d.next[firstChar] = &WordDictionary{next: make(map[byte]*WordDictionary)}
	}

	d.next[firstChar].AddWord(word[1:])
}

func (d *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return d.isEnd
	}

	firstChar := word[0]

	if firstChar == '.' {
		for _, child := range d.next {
			if child.Search(word[1:]) {
				return true
			}
		}
		return false
	}

	if d.next[firstChar] == nil {
		return false
	}

	return d.next[firstChar].Search(word[1:])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

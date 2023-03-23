// Source : https://leetcode.cn/problems/implement-trie-prefix-tree
// Date   : 2023-03-24

/**********************************************************************************
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：

	Trie() 初始化前缀树对象。
	void insert(String word) 向前缀树中插入字符串 word 。
	boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
	boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。


示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]
解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True


提示：

	1 <= word.length, prefix.length <= 2000
	word 和 prefix 仅由小写英文字母组成
	insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
**********************************************************************************/

package implementTriePrefixTree

type Trie struct {
	isEnd bool
	dict  map[byte]*Trie
}

func Constructor() Trie {
	return Trie{dict: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		t.isEnd = true
		return
	}

	if t.dict[word[0]] == nil {
		t.dict[word[0]] = &Trie{
			dict: make(map[byte]*Trie),
		}
	}
	t.dict[word[0]].Insert(word[1:])
}

func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return t.isEnd
	}

	if t.dict[word[0]] == nil {
		return false
	}

	return t.dict[word[0]].Search(word[1:])
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	if t.dict[prefix[0]] == nil {
		return false
	}

	return t.dict[prefix[0]].StartsWith(prefix[1:])
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

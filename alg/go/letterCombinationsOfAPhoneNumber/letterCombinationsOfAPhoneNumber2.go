package letterCombinations

/*
	这道题可以用回溯算法来解决。回溯算法是一种深度优先遍历算法的特例，它在遍历的过程中不断地尝试各种可能的解，并在不满足条件的情况下回溯（回退）到上一步继续搜索，直到找到合适的解为止。

	具体实现方法如下：

	1. 首先定义一个数组letterMap用来存储数字和字母的对应关系，然后定义一个结果集res用来存储所有的结果。

	2. 定义一个递归函数dfs，用来搜索所有可能的字母组合。函数的参数包括当前字符串s，当前遍历到的数字在原始字符串digits中的下标index，以及结果集res。

	3. 如果当前遍历的下标index等于原始字符串digits的长度，说明已经找到了一种字母组合，将其加入到结果集res中，然后直接返回。

	4. 否则，从letterMap中找到当前数字对应的字母集合，然后遍历这个集合中的每个字母。对于每个字母，将其加入到当前字符串s的末尾，
	   并以s和index+1作为参数调用递归函数dfs，在函数返回后，需要将s的末尾字母删除，以便遍历下一个字母。

	最后在主函数中调用递归函数dfs，并返回结果集res。
*/

func letterCombinations2(digits string) []string {
	letterMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	if len(digits) == 0 {
		return res
	}

	dfs(letterMap, digits, 0, "", &res)
	return res
}

func dfs(letterMap map[byte]string, digits string, index int, s string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, s)
		return
	}

	letters := letterMap[digits[index]]
	for i := 0; i < len(letters); i++ {
		s += string(letters[i])
		dfs(letterMap, digits, i+1, s, res)
		s = s[:len(s)-1]
	}
}

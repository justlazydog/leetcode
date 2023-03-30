package scrambleString

//这里，我们使用了一个三维数组mem来保存搜索的结果。
//mem[i][j][k]表示s1从下标i开始，s2从下标j开始，长度为k的子串是否可以通过扰乱得到。
//mem[i][j][k]的取值范围为{-1, 0, 1}，其中-1表示未搜索过，0表示无法通过扰乱得到，1表示可以通过扰乱得到。
//
//在每次搜索时，我们首先检查是否已经搜索过。
//如果已经搜索过，直接返回结果。
//接着，我们判断s1[i:i+k]和s2[j:j+k]是否相同，如果相同，直接返回1。
//如果不相同，检查它们是否由相同的字符组成。如果不是，返回0。
//如果是，则枚举分割点，判断两个子串是否可以通过扰乱得到，如果有一种方案可以，则返回1。
//如果没有，则返回0。
//
//这里需要注意的是，在判断两个子串是否可以通过扰乱得到时，我们需要递归调用dfs函数，并将其结果保存到mem数组中。这样可以避免重复计算，从而提高效率。
//
//最后，我们将dfs(s1, 0, s2, 0, n, mem)的返回值与1进行比较，如果相等，则s1可以通过扰乱得到s2，返回true；否则返回false。

func isScramble2(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	// mem[i][j][k]表示s1从i开始、s2从j开始、长度为k的子串是否可以通过扰乱得到
	mem := make([][][]int8, n)
	for i := range mem {
		mem[i] = make([][]int8, n)
		for j := range mem[i] {
			mem[i][j] = make([]int8, n+1)
		}
	}
	return dfs(s1, 0, s2, 0, n, mem) == 1
}

func dfs(s1 string, i int, s2 string, j int, k int, mem [][][]int8) int8 {
	if mem[i][j][k] != 0 {
		return mem[i][j][k]
	}
	// 判断s1[i:i+k]是否可以通过扰乱得到s2[j:j+k]
	if s1[i:i+k] == s2[j:j+k] {
		mem[i][j][k] = 1
		return 1
	}
	// 判断s1[i:i+k]和s2[j:j+k]是否由相同的字符组成
	if !checkIfSimilar(s1, i, s2, j, k) {
		mem[i][j][k] = -1
		return -1
	}
	// 枚举分割点
	for l := 1; l < k; l++ {
		// 检查s1[i:i+l]和s2[j:j+l]是否可以通过扰乱得到，s1[i+l:i+k]和s2[j+l:j+k]是否可以通过扰乱得到
		if dfs(s1, i, s2, j, l, mem) == 1 && dfs(s1, i+l, s2, j+l, k-l, mem) == 1 {
			mem[i][j][k] = 1
			return 1
		}
		// 检查s1[i:i+l]和s2[j+k-l:j+k]是否可以通过扰乱得到，s1[i+l:i+k]和s2[j:j+k-l]是否可以通过扰乱得到
		if dfs(s1, i, s2, j+k-l, l, mem) == 1 && dfs(s1, i+l, s2, j, k-l, mem) == 1 {
			mem[i][j][k] = 1
			return 1
		}
	}
	mem[i][j][k] = -1
	return -1
}

// 判断s1[i:i+k]和s2[j:j+k]是否由相同的字符组成
func checkIfSimilar(s1 string, i int, s2 string, j int, k int) bool {
	cnt := make([]int, 26)
	for p := 0; p < k; p++ {
		cnt[s1[i+p]-'a']++
		cnt[s2[j+p]-'a']--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

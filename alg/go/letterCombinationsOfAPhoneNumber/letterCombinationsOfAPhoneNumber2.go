package letterCombinations

func letterCombinations2(digits string) []string {
	dic := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}

	var dfs func(index int, temp string)
	dfs = func(index int, temp string) {
		if index == len(digits) {
			res = append(res, temp)
			return
		}

		word := dic[digits[index]]
		for _, val := range word {
			temp += string(val)
			dfs(index+1, temp)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, "")
	return res
}

// Source : https://leetcode.cn/problems/permutations
// Date   : 2023-03-09

/**********************************************************************************
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：

输入：nums = [1]
输出：[[1]]


提示：

	1 <= nums.length <= 6
	-10 <= nums[i] <= 10
	nums 中的所有整数 互不相同
**********************************************************************************/

package permutations

/*
题目要求得到给定数组的全排列，使用回溯算法可以轻松解决这个问题。回溯算法的核心思想是枚举所有可能的情况，当满足特定条件时回溯并继续枚举。

在这个问题中，我们可以定义一个 backtrack 函数来查找所有可能的排列。在 backtrack 函数中，我们维护一个 path 数组，用于存储当前已经选择的数字。当 path 中的数字数量与 nums 数组中的数字数量相同时，表示已经找到了一种可能的排列。此时，我们将 path 数组复制一份并添加到结果数组 res 中。

在每次递归中，我们循环遍历 nums 数组中的所有数字。如果该数字已经在 path 数组中出现过，则跳过该数字；否则，将该数字添加到 path 数组中，并递归调用 backtrack 函数。递归结束后，需要将 path 数组的最后一个数字删除，以便在下一次循环中选择其他数字。

最后，我们调用 permute 函数并将 nums 数组和一个空的 path 数组作为参数传递。在 permute 函数中，我们初始化结果数组 res，然后调用 backtrack 函数查找所有可能的排列。最终，我们返回结果数组 res。
*/

func permute(nums []int) [][]int {
	var res [][]int
	backtrack(nums, []int{}, &res)
	return res
}

func backtrack(nums []int, path []int, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for _, num := range nums {
		if contains(path, num) {
			continue
		}
		path = append(path, num)
		backtrack(nums, path, res)
		path = path[:len(path)-1]
	}
}

func contains(arr []int, val int) bool {
	for _, num := range arr {
		if num == val {
			return true
		}
	}
	return false
}

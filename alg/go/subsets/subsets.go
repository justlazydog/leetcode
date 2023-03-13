// Source : https://leetcode.cn/problems/subsets
// Date   : 2023-03-13

/**********************************************************************************
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

	1 <= nums.length <= 10
	-10 <= nums[i] <= 10
	nums 中的所有元素 互不相同
**********************************************************************************/

package subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	backtrack(0, path, nums, &res)
	return res
}

func backtrack(start int, path []int, nums []int, res *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrack(i+1, path, nums, res)
		path = path[:len(path)-1]
	}
}

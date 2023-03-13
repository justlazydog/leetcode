// Source : https://leetcode.cn/problems/merge-intervals
// Date   : 2023-03-13

/**********************************************************************************
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

	1 <= intervals.length <= 104
	intervals[i].length == 2
	0 <= starti <= endi <= 104
**********************************************************************************/

package mergeIntervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	return f(intervals[0], 1, intervals)
}

func f(start []int, i int, nums [][]int) [][]int {
	var res [][]int
	for ; i < len(nums); i++ {
		if start[1] >= nums[i][1] {
			continue
		} else if start[1] >= nums[i][0] && start[1] < nums[i][1] {
			start[1] = nums[i][1]
			return append(res, f(start, i+1, nums)...)
		} else {
			res = append(res, start)
			return append(res, f(nums[i], i+1, nums)...)
		}
	}

	if i == len(nums) {
		res = append(res, start)
	}
	return res
}

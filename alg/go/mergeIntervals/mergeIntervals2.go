package mergeIntervals

import "sort"

func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 先按区间左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			// 当前区间和上一个区间重叠，合并
			last[1] = max(last[1], curr[1])
		} else {
			// 当前区间和上一个区间不重叠，加入合并后的区间集合中
			merged = append(merged, curr)
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

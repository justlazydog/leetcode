// Source : https://leetcode.cn/problems/largest-rectangle-in-histogram
// Date   : 2023-03-13

/**********************************************************************************
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:


输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10

示例 2：


输入： heights = [2,4]
输出： 4

提示：

	1 <= heights.length <=105
	0 <= heights[i] <= 104
**********************************************************************************/

package largestRectangleInHistogram

/*
	使用单调栈来解决。

	具体来说，我们可以遍历整个数组，维护一个单调递增的栈，栈中保存的是数组中每个元素的下标。
	如果当前元素的高度比栈顶元素的高度低，那么我们就可以弹出栈顶元素，计算以栈顶元素为高度的最大矩形的面积。
	在计算过程中，以栈顶元素的下一个元素的下标为右边界，以栈底元素的下标为左边界，计算矩形的面积。
	每次计算完以栈顶元素为高度的最大矩形面积后，将该元素的下标入栈。
*/

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	ans := 0
	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			ans = max(ans, height*width)
		}
		stack = append(stack, i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

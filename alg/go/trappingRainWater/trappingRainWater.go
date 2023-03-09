// Source : https://leetcode.cn/problems/trapping-rain-water
// Date   : 2023-03-09

/**********************************************************************************
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：


输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：

输入：height = [4,2,0,3,2,5]
输出：9


提示：

	n == height.length
	1 <= n <= 2 * 104
	0 <= height[i] <= 105
**********************************************************************************/

package trappingRainWater

/*
	这个解法使用了双指针法，通过两个指针从两端向中间扫描来计算接雨水的量。我们首先初始化两个指针 left 和 right，分别指向数组的最左端和最右端。
	同时，我们初始化两个变量 leftMax 和 rightMax，分别表示 left 左侧和 right 右侧的最大高度。

	接下来，我们使用一个 while 循环，不断地将指针向中间移动。在每一步中，我们需要比较 height[left] 和 height[right]，
	将较小的一侧向中间移动，并更新 leftMax 和 rightMax。
	如果移动过程中发现当前位置的高度小于 leftMax 或 rightMax，说明当前位置可以存储雨水，需要将雨水量加入到答案中。

	最后，当 left 和 right 相遇时，我们就可以得到接雨水的总量，返回答案即可。
*/

func trap(height []int) int {
	var res int
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] < leftMax {
				res += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		} else {
			right--
			if height[right] < rightMax {
				res += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		}
	}
	return res
}

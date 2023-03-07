// Source : https://leetcode.cn/problems/4sum
// Date   : 2023-03-07

/**********************************************************************************
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

	0 <= a, b, c, d < n
	a、b、c 和 d 互不相同
	nums[a] + nums[b] + nums[c] + nums[d] == target

你可以按 任意顺序 返回答案 。

示例 1：

输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：

输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]


提示：

	1 <= nums.length <= 200
	-109 <= nums[i] <= 109
	-109 <= target <= 109
**********************************************************************************/

package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			sum := nums[i] + nums[j]
			diff, left, right := target-sum, j+1, len(nums)-1
			for left < right {
				sum = nums[left] + nums[right]
				if sum == diff {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--

					for ; left < right && nums[left] == nums[left-1]; left++ {
					}

					for ; left < right && nums[right] == nums[right+1]; right-- {
					}
				}

				if sum > diff {
					right--
				}

				if sum < diff {
					left++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

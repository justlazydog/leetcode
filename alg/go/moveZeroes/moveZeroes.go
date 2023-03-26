// Source : https://leetcode.cn/problems/move-zeroes
// Date   : 2023-03-26

/**********************************************************************************
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:

输入: nums = [0]
输出: [0]

提示:

	1 <= nums.length <= 104
	-231 <= nums[i] <= 231 - 1


进阶：你能尽量减少完成的操作次数吗？
**********************************************************************************/

package moveZeroes

//这道题可以用两个指针 left 和 right 来解决，初始时 left 和 right 都指向数组的第一个元素。
//我们让 left 和 right 向右移动，如果 left 指向的元素是 0，而 right 指向的元素不是 0，则交换 left 和 right 指向的元素，并将 left 向右移动一位。

func moveZeroes(nums []int) {
	n := len(nums)
	left, right := 0, 0
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

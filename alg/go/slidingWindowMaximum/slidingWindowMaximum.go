// Source : https://leetcode.cn/problems/sliding-window-maximum
// Date   : 2023-03-25

/**********************************************************************************
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。

示例 1：

输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：

输入：nums = [1], k = 1
输出：[1]


提示：

	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
	1 <= k <= nums.length
**********************************************************************************/

package slidingWindowMaximum

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	// 用双端队列（deque）存储可能成为滑动窗口最大值的数的下标
	deque := make([]int, 0)

	// 初始化 deque，保证 deque[0] 是当前窗口中的最大值的下标
	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	n := len(nums)
	res := make([]int, n-k+1)
	res[0] = nums[deque[0]]

	// 遍历 nums，维护 deque，更新 res
	for i := k; i < n; i++ {
		// 如果 deque[0] 已经超出当前窗口，将其从 deque 中删除
		if deque[0] <= i-k {
			deque = deque[1:]
		}

		// 删除 deque 中所有小于 nums[i] 的数的下标
		for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		res[i-k+1] = nums[deque[0]]
	}

	return res
}

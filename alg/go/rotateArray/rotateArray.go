// Source : https://leetcode.cn/problems/rotate-array
// Date   : 2023-03-22

/**********************************************************************************
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：

	1 <= nums.length <= 105
	-231 <= nums[i] <= 231 - 1
	0 <= k <= 105


进阶：

	尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
	你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
**********************************************************************************/

package rotateArray

func rotate(nums []int, k int) {
	n := len(nums)
	if k >= n {
		k = k % n
	}
	tmp := append(nums[n-k:], nums[:n-k]...)
	for i, num := range tmp {
		nums[i] = num
	}
}

// 暴力移动
// 最简单的方法是使用一个暴力的循环来移动数组中的元素，每次将最后一个元素插入到数组的最前面，重复k次。
// 这个方法的时间复杂度是O(n*k)，其中n是数组的长度。
func rotateA(nums []int, k int) {
	k = k % len(nums)
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// 翻转数组
// 更加高效的方法是将数组翻转，然后将前k个元素翻转，再将剩下的元素翻转。
// 这个方法的时间复杂度是O(n)，其中n是数组的长度。以下是代码实现：
func rotateB(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 环状替换
// 我们可以将数组中的元素看做是一个环，我们从一个元素开始，将这个元素向右移动k个位置，
// 然后再将这个位置的元素向右移动k个位置，一直循环下去，直到所有的元素都移动了。
// 这个方法需要处理一些边界条件，但是代码实现也很简单
func rotateC(nums []int, k int) {
	k = k % len(nums)
	count := 0
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

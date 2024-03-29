// Source : https://leetcode.cn/problems/majority-element
// Date   : 2023-03-22

/**********************************************************************************
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2


提示：

	n == nums.length
	1 <= n <= 5 * 104
	-109 <= nums[i] <= 109


进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
**********************************************************************************/

package majorityElement

//该算法的思路是使用“摩尔投票法”，遍历整个数组，如果当前数等于候选人，则将计数器加1，否则将计数器减1。
//如果计数器为0，则将当前数设置为候选人。由于出现次数超过一半的元素最多只有一个，所以最终的候选人即为众数。
//摩尔投票法的应用非常广泛，不仅仅局限于求出现次数超过一半的元素，还可以求出现次数超过1/k的元素（k>2），只需要维护k-1个计数器和k-1个候选人即可。

func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		// 如果count为0，则将当前数设置为候选人
		if count == 0 {
			candidate = num
		}
		// 如果当前数等于候选人，则count加1，否则count减1
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

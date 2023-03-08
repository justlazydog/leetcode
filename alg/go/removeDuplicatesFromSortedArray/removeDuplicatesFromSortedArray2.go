package removeDuplicates

/*
	在该解法中，我们使用双指针来遍历数组。
	i表示新数组的当前长度，j表示原数组的当前位置。
	当发现一个新的不同元素时，将其添加到新数组中，并更新指针i。
	最后，返回新数组的长度即可。
*/

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

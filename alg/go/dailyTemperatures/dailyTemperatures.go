// Source : https://leetcode.cn/problems/daily-temperatures
// Date   : 2023-03-29

/**********************************************************************************
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]

示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]

示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]

提示：

	1 <= temperatures.length <= 105
	30 <= temperatures[i] <= 100
**********************************************************************************/

package dailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stack []int

	for i := 0; i < n; i++ {
		// 如果当前元素比栈顶元素大，则栈顶元素的下一个更大元素即为当前元素
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[j] = i - j
		}
		// 把当前元素的下标压入栈中
		stack = append(stack, i)
	}

	return res
}

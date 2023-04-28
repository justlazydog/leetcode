// Source : https://leetcode.cn/problems/happy-number
// Date   : 2023-04-28

/**********************************************************************************
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」 定义为：

	对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
	然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
	如果这个过程 结果为 1，那么这个数就是快乐数。

如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

示例 2：

输入：n = 2
输出：false


提示：

	1 <= n <= 231 - 1
**********************************************************************************/

package happyNumber

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}

func squareSum(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n = n / 10
	}
	return sum
}

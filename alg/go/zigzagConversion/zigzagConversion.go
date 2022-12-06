// Source : https://leetcode.cn/problems/zigzag-conversion
// Date   : 2022-12-06

/**********************************************************************************
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：

输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
示例 2：

输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"
解释：
P     I    N
A   L S  I G
Y A   H R
P     I

示例 3：

输入：s = "A", numRows = 1
输出："A"


提示：

	1 <= s.length <= 1000
	s 由英文字母（小写和大写）、',' 和 '.' 组成
	1 <= numRows <= 1000
**********************************************************************************/

package zigzagConversion

func convert(s string, numRows int) string {
	// special case just return
	if numRows == 1 || len(s) < numRows {
		return s
	}

	str := ""
	for row := 1; row <= numRows; row++ {
		i := row - 1  // s'index
		flag := false // flag means whether s needs to be turned

		for i < len(s) {
			str += string(s[i])

			// row one and last row don't need to be turned
			if row == 1 || row == numRows {
				i += 2 * (numRows - 1)
			} else {
				// other rows index rules
				if !flag {
					i += 2 * (numRows - row)
				} else {
					i += 2 * (numRows - (numRows - row + 1))
				}

				if i < len(s) {
					flag = !flag
				}
			}
		}
	}
	return str
}

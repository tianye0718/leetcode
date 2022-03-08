/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (22.13%)
 * Likes:    851
 * Dislikes: 0
 * Total Accepted:    150.1K
 * Total Submissions: 678.2K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 *
 * 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
 * -2
 *
 *
 *
 * 示例 1:
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
 *
 * 示例 2:
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = truncate(-2.33333..) = -2
 *
 *
 *
 * 提示：
 *
 *
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 *
 *
 */
package main

import "math"

// @lc code=start
func divide(dividend int, divisor int) (res int) {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// 获取结果的正负号
	negative := (dividend > 0) != (divisor > 0)
	// 取绝对值，要考虑溢出的情况，用无符号整数来储存结果
	a, b := abs(dividend), abs(divisor)

	for i := 31; i >= 0; i-- {
		if a>>i >= b { // 找出最大的可用2的次幂表示的商
			res += 1 << i
			a -= b << i
		}
	}
	if negative {
		return -res
	}
	return res
}
func abs(x int) (res uint) {
	if x < 0 {
		res = uint(-(x + 1))
		res = res + 1
	} else {
		res = uint(x)
	}
	return res
}

// @lc code=end

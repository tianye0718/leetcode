/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (37.82%)
 * Likes:    886
 * Dislikes: 0
 * Total Accepted:    261.7K
 * Total Submissions: 692K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n^ ）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 *
 *
 * 示例 3：
 *
 *
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 *
 * 提示：
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * -10^4 <= x^n <= 10^4
 *
 *
 */
package main

// @lc code=start
func myPow(x float64, n int) float64 {
	// 3^10, 10的二进制表示是1010
	// 3^10 可以表示为3^8 * 3^0 * 3^2 * 3^0, 3^0 省略
	res := 1.0
	// 取绝对值
	var exp uint
	if n < 0 {
		exp = uint(-(n + 1))
		exp = exp + 1
	} else {
		exp = uint(n)
	}

	for exp != 0 {
		if exp&1 == 1 { // 最后一位为1，则结果累乘
			res *= x
		}
		x *= x // 被乘数累乘
		exp >>= 1
	}
	if n < 0 {
		res = 1 / res
	}
	return res
}

// @lc code=end

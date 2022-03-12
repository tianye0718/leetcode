/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (70.02%)
 * Likes:    1606
 * Dislikes: 0
 * Total Accepted:    211.6K
 * Total Submissions: 302.1K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：5
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package leetcode0096

// @lc code=start
var memo []int

func numTrees(n int) int {
	memo = make([]int, n+1)
	return count(1, n)
}
func count(lo int, hi int) int {
	if lo > hi {
		return 1
	}
	if memo[hi-lo] != 0 {
		return memo[hi-lo]
	}
	res := 0
	for root := lo; root <= hi; root++ {
		left := count(lo, root-1)
		right := count(root+1, hi)
		res += left * right
	}
	memo[hi-lo] = res
	return res
}

// @lc code=end

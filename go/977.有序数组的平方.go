/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 *
 * https://leetcode-cn.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (69.58%)
 * Likes:    434
 * Dislikes: 0
 * Total Accepted:    255.9K
 * Total Submissions: 367.9K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-4,-1,0,3,10]
 * 输出：[0,1,9,16,100]
 * 解释：平方后，数组变为 [16,1,0,9,100]
 * 排序后，数组变为 [0,1,9,16,100]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [-7,-3,2,3,11]
 * 输出：[4,9,9,49,121]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 10^4
 * -10^4
 * nums 已按 非递减顺序 排序
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 请你设计时间复杂度为 O(n) 的算法解决本问题
 *
 *
 */
package main

// @lc code=start
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	index := len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; {
		if square(nums[i]) > square(nums[j]) {
			res[index] = square(nums[i])
			i++
		} else {
			res[index] = square(nums[j])
			j--
		}
		index--
	}
	return res
}
func square(x int) int {
	return x * x
}

// @lc code=end

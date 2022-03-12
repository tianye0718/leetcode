/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (55.21%)
 * Likes:    4355
 * Dislikes: 0
 * Total Accepted:    852K
 * Total Submissions: 1.5M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 子数组 是数组中的一个连续部分。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [5,4,-1,7,8]
 * 输出：23
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 *
 */
package leetcode0053

// @lc code=start
// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	maxSum := nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(dp[i-1]+nums[i], nums[i])
// 		if dp[i] > maxSum {
// 			maxSum = dp[i]
// 		}
// 	}
// 	return maxSum
// }

func maxSubArray(nums []int) int {
	prev := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		prev = max(prev+nums[i], nums[i])
		if prev > maxSum {
			maxSum = prev
		}
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (45.07%)
 * Likes:    1752
 * Dislikes: 0
 * Total Accepted:    375.5K
 * Total Submissions: 831.9K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
 *
 * 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
 *
 * 你可以认为每种硬币的数量是无限的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3
 * 解释：11 = 5 + 5 + 1
 *
 * 示例 2：
 *
 *
 * 输入：coins = [2], amount = 3
 * 输出：-1
 *
 * 示例 3：
 *
 *
 * 输入：coins = [1], amount = 0
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */
package main

// @lc code=start
// recursive
// var memo []int

// func coinChange(coins []int, amount int) int {
// 	for i := 0; i < amount+1; i++ {
// 		memo = append(memo, -666)
// 	}
// 	return dp(coins, amount)
// }
// func dp(coins []int, amount int) int {
// 	if amount < 0 {
// 		return -1
// 	}
// 	if amount == 0 {
// 		return 0
// 	}
// 	if memo[amount] != -666 {
// 		return memo[amount]
// 	}
// 	res := math.MaxInt32
// 	for _, coin := range coins {
// 		subProblem := dp(coins, amount-coin)
// 		if subProblem == -1 {
// 			continue
// 		}
// 		res = min(res, subProblem+1)
// 	}
// 	if res == math.MaxInt32 {
// 		memo[amount] = -1
// 	} else {
// 		memo[amount] = res
// 	}
// 	return memo[amount]
// }
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}

// @lc code=end

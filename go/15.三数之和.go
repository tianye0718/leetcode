/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (34.55%)
 * Likes:    4384
 * Dislikes: 0
 * Total Accepted:    823.3K
 * Total Submissions: 2.4M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */
package main

import "sort"

// @lc code=start

// Two pointers

// func threeSum(nums []int) (res [][]int) {
// 	sort.Ints(nums)
// 	for i := 0; i < len(nums); i++ {
// 		n1 := nums[i]
// 		if n1 > 0 {
// 			return res
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
// 		l, r := i+1, len(nums)-1
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
// 			if n1+n2+n3 == 0 {
// 				res = append(res, []int{n1, n2, n3})
// 				// remove duplicate
// 				for l < r && n2 == nums[l] {
// 					l++
// 				}
// 				for l < r && n3 == nums[r] {
// 					r--
// 				}
// 			} else if n1+n2+n3 < 0 {
// 				l++
// 			} else {
// 				r--
// 			}
// 		}

// 	}
// 	return res
// }

// Backtracking
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 3, 0, 0)
}

func nSum(nums []int, n int, start int, target int) (res [][]int) {
	sz := len(nums)
	if n < 2 || n > sz {
		return res
	}
	if n > 2 {
		for i := start; i < sz; i++ {
			sub := nSum(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	} else if n == 2 {
		lo, hi := start, sz-1
		for lo < hi {
			left := nums[lo]
			right := nums[hi]
			if left+right > target {
				hi--
			} else if left+right < target {
				lo++
			} else {
				res = append(res, []int{left, right})
				for lo < hi && left == nums[lo] {
					lo++
				}
				for lo < hi && right == nums[hi] {
					hi--
				}
			}
		}
	}
	return res
}

// @lc code=end

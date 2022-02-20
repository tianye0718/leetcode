/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		j, ok := m[target-num]
		if ok {
			return []int{i, j}
		} else {
			m[num] = i
		}
	}
	return []int{}
}

// @lc code=end

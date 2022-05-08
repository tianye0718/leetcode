/*
 * @lc app=leetcode id=1512 lang=golang
 *
 * [1512] Number of Good Pairs
 *
 * https://leetcode.com/problems/number-of-good-pairs/description/
 *
 * algorithms
 * Easy (87.98%)
 * Likes:    2441
 * Dislikes: 141
 * Total Accepted:    292.5K
 * Total Submissions: 332.4K
 * Testcase Example:  '[1,2,3,1,1,3]'
 *
 * Given an array of integers nums, return the number of good pairs.
 *
 * A pair (i, j) is called good if nums[i] == nums[j] and i < j.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1,1,3]
 * Output: 4
 * Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,1,1,1]
 * Output: 6
 * Explanation: Each pair in the array are good.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,2,3]
 * Output: 0
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 *
 *
 */
package leetcode1512

// @lc code=start
func numIdenticalPairs(nums []int) (ans int) {
	if len(nums) < 2 {
		return
	}
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		if m[num] > 0 {
			ans += m[num]
		}
		m[num]++
	}
	return
}

// @lc code=end

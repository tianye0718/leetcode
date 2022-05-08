/*
 * @lc app=leetcode id=1052 lang=golang
 *
 * [1052] Grumpy Bookstore Owner
 *
 * https://leetcode.com/problems/grumpy-bookstore-owner/description/
 *
 * algorithms
 * Medium (56.72%)
 * Likes:    1138
 * Dislikes: 95
 * Total Accepted:    52.9K
 * Total Submissions: 93.3K
 * Testcase Example:  '[1,0,1,2,1,1,7,5]\n[0,1,0,1,0,1,0,1]\n3'
 *
 * There is a bookstore owner that has a store open for n minutes. Every
 * minute, some number of customers enter the store. You are given an integer
 * array customers of length n where customers[i] is the number of the customer
 * that enters the store at the start of the i^th minute and all those
 * customers leave after the end of that minute.
 *
 * On some minutes, the bookstore owner is grumpy. You are given a binary array
 * grumpy where grumpy[i] is 1 if the bookstore owner is grumpy during the i^th
 * minute, and is 0 otherwise.
 *
 * When the bookstore owner is grumpy, the customers of that minute are not
 * satisfied, otherwise, they are satisfied.
 *
 * The bookstore owner knows a secret technique to keep themselves not grumpy
 * for minutes consecutive minutes, but can only use it once.
 *
 * Return the maximum number of customers that can be satisfied throughout the
 * day.
 *
 *
 * Example 1:
 *
 *
 * Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], minutes =
 * 3
 * Output: 16
 * Explanation: The bookstore owner keeps themselves not grumpy for the last 3
 * minutes.
 * The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 +
 * 5 = 16.
 *
 *
 * Example 2:
 *
 *
 * Input: customers = [1], grumpy = [0], minutes = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * n == customers.length == grumpy.length
 * 1 <= minutes <= n <= 2 * 10^4
 * 0 <= customers[i] <= 1000
 * grumpy[i] is either 0 or 1.
 *
 *
 */
package leetcode1052

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {
	size := len(customers)
	preSumWithoutGrumpy := make([]int, size+1)
	preSumWithGrumpy := make([]int, size+1)

	preSumWithoutGrumpy[0], preSumWithGrumpy[0] = 0, 0
	for i := 1; i <= size; i++ {
		preSumWithoutGrumpy[i] = preSumWithoutGrumpy[i-1] + customers[i-1]
		preSumWithGrumpy[i] = preSumWithGrumpy[i-1] + customers[i-1]*((grumpy[i-1]+1)%2)
	}

	for i, j := 0, minutes-1; j < size; i, j = i+1, j+1 {
		currWindowWithGrumpy := preSumWithGrumpy[j+1] - preSumWithGrumpy[i]
		currWindowWithoutGrumpy := preSumWithoutGrumpy[j+1] - preSumWithoutGrumpy[i]
		sum := preSumWithGrumpy[size] + currWindowWithoutGrumpy - currWindowWithGrumpy
		if sum > ans {
			ans = sum
		}
	}
	return
}

// @lc code=end

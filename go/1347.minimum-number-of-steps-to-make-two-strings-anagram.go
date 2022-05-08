/*
 * @lc app=leetcode id=1347 lang=golang
 *
 * [1347] Minimum Number of Steps to Make Two Strings Anagram
 *
 * https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description/
 *
 * algorithms
 * Medium (76.40%)
 * Likes:    1221
 * Dislikes: 57
 * Total Accepted:    89.9K
 * Total Submissions: 117.6K
 * Testcase Example:  '"bab"\n"aba"'
 *
 * You are given two strings of the same length s and t. In one step you can
 * choose any character of t and replace it with another character.
 *
 * Return the minimum number of steps to make t an anagram of s.
 *
 * An Anagram of a string is a string that contains the same characters with a
 * different (or the same) ordering.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "bab", t = "aba"
 * Output: 1
 * Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram
 * of s.
 *
 *
 * Example 2:
 *
 *
 * Input: s = "leetcode", t = "practice"
 * Output: 5
 * Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper
 * characters to make t anagram of s.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "anagram", t = "mangaar"
 * Output: 0
 * Explanation: "anagram" and "mangaar" are anagrams.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * s.length == t.length
 * s and t consist of lowercase English letters only.
 *
 *
 */
package leetcode1347

// @lc code=start
func minSteps(s string, t string) (ans int) {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		if m[c] == 0 {
			ans++
		} else {
			m[c]--
		}
	}
	return
}

// @lc code=end

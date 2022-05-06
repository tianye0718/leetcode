/*
 * @lc app=leetcode id=1876 lang=golang
 *
 * [1876] Substrings of Size Three with Distinct Characters
 *
 * https://leetcode.com/problems/substrings-of-size-three-with-distinct-characters/description/
 *
 * algorithms
 * Easy (69.67%)
 * Likes:    572
 * Dislikes: 19
 * Total Accepted:    41.8K
 * Total Submissions: 59.9K
 * Testcase Example:  '"xyzzaz"'
 *
 * A string is good if there are no repeated characters.
 *
 * Given a string s​​​​​, return the number of good substrings of length three
 * in s​​​​​​.
 *
 * Note that if there are multiple occurrences of the same substring, every
 * occurrence should be counted.
 *
 * A substring is a contiguous sequence of characters in a string.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "xyzzaz"
 * Output: 1
 * Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and
 * "zaz".
 * The only good substring of length 3 is "xyz".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "aababcabc"
 * Output: 4
 * Explanation: There are 7 substrings of size 3: "aab", "aba", "bab", "abc",
 * "bca", "cab", and "abc".
 * The good substrings are "abc", "bca", "cab", and "abc".
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 100
 * s​​​​​​ consists of lowercase English letters.
 *
 *
 */
package leetcode1876

// @lc code=start
func countGoodSubstrings(s string) (cnt int) {
	size := len(s)
	if size < 3 {
		return 0
	}
	m := make(map[byte]int)
	i, j := 0, 0
	for j < size {
		m[s[j]]++
		for m[s[j]] > 1 {
			m[s[i]]--
			i++
		}
		if j-i == 2 {
			cnt++
			m[s[i]]--
			i++
		}
		j++
	}
	return
}

// @lc code=end

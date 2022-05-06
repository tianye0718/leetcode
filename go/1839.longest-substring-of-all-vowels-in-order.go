/*
 * @lc app=leetcode id=1839 lang=golang
 *
 * [1839] Longest Substring Of All Vowels in Order
 *
 * https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/description/
 *
 * algorithms
 * Medium (47.93%)
 * Likes:    400
 * Dislikes: 13
 * Total Accepted:    18.3K
 * Total Submissions: 38.2K
 * Testcase Example:  '"aeiaaioaaaaeiiiiouuuooaauuaeiu"'
 *
 * A string is considered beautiful if it satisfies the following
 * conditions:
 *
 *
 * Each of the 5 English vowels ('a', 'e', 'i', 'o', 'u') must appear at least
 * once in it.
 * The letters must be sorted in alphabetical order (i.e. all 'a's before 'e's,
 * all 'e's before 'i's, etc.).
 *
 *
 * For example, strings "aeiou" and "aaaaaaeiiiioou" are considered beautiful,
 * but "uaeio", "aeoiu", and "aaaeeeooo" are not beautiful.
 *
 * Given a string word consisting of English vowels, return the length of the
 * longest beautiful substring of word. If no such substring exists, return 0.
 *
 * A substring is a contiguous sequence of characters in a string.
 *
 *
 * Example 1:
 *
 *
 * Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
 * Output: 13
 * Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of
 * length 13.
 *
 * Example 2:
 *
 *
 * Input: word = "aeeeiiiioooauuuaeiou"
 * Output: 5
 * Explanation: The longest beautiful substring in word is "aeiou" of length
 * 5.
 *
 *
 * Example 3:
 *
 *
 * Input: word = "a"
 * Output: 0
 * Explanation: There is no beautiful substring, so return 0.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= word.length <= 5 * 10^5
 * word consists of characters 'a', 'e', 'i', 'o', and 'u'.
 *
 *
 */
package leetcode1839

// @lc code=start
func longestBeautifulSubstring(word string) int {
	size := len(word)
	if size < 5 {
		return 0
	}
	max := 0
	for i, j := 0, 1; j < size; j++ {
		if toVowelsIndex(word[j])-toVowelsIndex(word[j-1]) != 0 && toVowelsIndex(word[j])-toVowelsIndex(word[j-1]) != 1 {
			i = j
			continue
		}

		if word[i] != byte('a') {
			i++
			continue
		}
		if word[j] == byte('u') {
			if j-i+1 > max {
				max = j - i + 1
			}
		}
	}
	return max
}
func toVowelsIndex(v byte) byte {
	switch v {
	case byte('a'):
		return 1
	case byte('e'):
		return 2
	case byte('i'):
		return 3
	case byte('o'):
		return 4
	case byte('u'):
		return 5
	}
	return 0
}

// @lc code=end

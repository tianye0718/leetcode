/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (64.97%)
 * Likes:    507
 * Dislikes: 0
 * Total Accepted:    361.8K
 * Total Submissions: 556.6K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * s 和 t 仅包含小写字母
 *
 *
 *
 *
 * 进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */
package main

// @lc code=start
// use array

// func isAnagram(s string, t string) bool {
// 	records := [26]int{}
// 	for _, r := range s {
// 		records[r-rune('a')]++
// 	}
// 	for _, r := range t {
// 		records[r-rune('a')]--
// 	}
// 	return records == [26]int{}
// }

// use slice map
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		if v, ok := records[s[i]]; ok {
			records[s[i]] = v + 1
		} else {
			records[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if v, ok := records[t[i]]; ok && v >= 1 {
			records[t[i]] = v - 1
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

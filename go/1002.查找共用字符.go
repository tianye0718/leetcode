/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找共用字符
 *
 * https://leetcode-cn.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (72.75%)
 * Likes:    260
 * Dislikes: 0
 * Total Accepted:    62.4K
 * Total Submissions: 85.9K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * 给你一个字符串数组 words ，请你找出所有在 words 的每个字符串中都出现的共用字符（ 包括重复字符），并以数组形式返回。你可以按 任意顺序
 * 返回答案。
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["bella","label","roller"]
 * 输出：["e","l","l"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：words = ["cool","lock","cook"]
 * 输出：["c","o"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 100
 * 1 <= words[i].length <= 100
 * words[i] 由小写英文字母组成
 *
 *
 */
package main

// @lc code=start
func commonChars(words []string) (res []string) {
	benchmark := [26]int{}
	for _, c := range words[0] {
		benchmark[c-97]++
	}
	for i := 1; i < len(words); i++ {
		competor := [26]int{}
		for _, c := range words[i] {
			competor[c-97]++
		}
		for j := 0; j < 26; j++ {
			benchmark[j] = min(benchmark[j], competor[j])
		}
	}
	for i := 0; i < 26; i++ {
		for benchmark[i] != 0 {
			res = append(res, string(rune(i+'a')))
			benchmark[i]--
		}
	}
	return res
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 *
 * https://leetcode-cn.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (53.01%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    83.8K
 * Total Submissions: 158.1K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 *
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 *
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 *
 * 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 *
 *
 * 示例 2:
 *
 *
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：把最后一位反向旋转一次即可 "0000" -> "0009"。
 *
 *
 * 示例 3:
 *
 *
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：无法旋转到目标数字且不被锁定。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= deadends.length <= 500
 * deadends[i].length == 4
 * target.length == 4
 * target 不在 deadends 之中
 * target 和 deadends[i] 仅由若干位数字组成
 *
 *
 */
package leetcode0752

// @lc code=start
func spinUp(code string, i int) string {
	r := []rune(code)
	if r[i] == '9' {
		r[i] = '0'
	} else {
		r[i]++
	}
	return string(r)
}
func spinDown(code string, i int) string {
	r := []rune(code)
	if r[i] == '0' {
		r[i] = '9'
	} else {
		r[i]--
	}
	return string(r)
}

//////////////////BFS//////////////////////
// func openLock(deadends []string, target string) int {
// 	d := make(map[string]struct{}, 0)
// 	for _, v := range deadends {
// 		d[v] = struct{}{}
// 	}
// 	start := "0000"
// 	if _, ok := d[start]; ok {
// 		return -1
// 	}
// 	v := make(map[string]struct{}, 0)

// 	steps := 0
// 	q := []string{start}
// 	for len(q) > 0 {
// 		sz := len(q)
// 		for i := 0; i < sz; i++ {
// 			cur := q[0]
// 			q = q[1:]
// 			if cur == target {
// 				return steps
// 			}
// 			if _, ok := d[cur]; ok {
// 				continue
// 			}
// 			v[cur] = struct{}{}
// 			for j := 0; j < 4; j++ {
// 				up := spinUp(cur, j)
// 				if _, ok := v[up]; !ok {
// 					q = append(q, up)
// 				}
// 				down := spinDown(cur, j)
// 				if _, ok := v[down]; !ok {
// 					q = append(q, down)
// 				}
// 			}
// 		}
// 		steps++
// 	}
// 	return -1
// }

///////////////////////Bi-Direction BFS//////////////
func openLock(deadends []string, target string) int {
	start := "0000"
	deads := make(map[string]struct{}, 0)
	for _, v := range deadends {
		deads[v] = struct{}{}
	}
	if _, ok := deads[start]; ok {
		return -1
	}
	visited := make(map[string]struct{}, 0)
	steps := 0

	q1 := make(map[string]struct{}, 0)
	q2 := make(map[string]struct{}, 0)
	q1[start] = struct{}{}
	q2[target] = struct{}{}
	for len(q1) != 0 && len(q2) != 0 {
		// Small trick: always spread the smaller set
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		temp := make(map[string]struct{}, 0)
		for key := range q1 {
			if _, ok := q2[key]; ok {
				return steps
			}
			if _, ok := deads[key]; ok {
				continue
			}
			visited[key] = struct{}{}
			for i := 0; i < 4; i++ {
				up := spinUp(key, i)
				if _, ok := visited[up]; !ok {
					temp[up] = struct{}{}
				}
				down := spinDown(key, i)
				if _, ok := visited[down]; !ok {
					temp[down] = struct{}{}
				}
			}
		}
		q1 = q2
		q2 = temp
		steps++
	}
	return -1
}

// @lc code=end

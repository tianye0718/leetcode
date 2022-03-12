/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (48.42%)
 * Likes:    992
 * Dislikes: 0
 * Total Accepted:    224.1K
 * Total Submissions: 462.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -100
 *
 *
 */
package leetcode0054

// @lc code=start
func spiralOrder(matrix [][]int) (res []int) {
	m, n := len(matrix), len(matrix[0])
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	status := make([][]bool, m)
	for i := range status {
		status[i] = make([]bool, n)
	}
	x, y, d := 0, 0, 0

	for i := 0; i < m*n; i++ {
		res = append(res, matrix[x][y])
		status[x][y] = true
		a, b := x+dx[d], y+dy[d]
		if a < 0 || a >= m || b < 0 || b >= n || status[a][b] {
			d = (d + 1) % 4
		}
		x = x + dx[d]
		y = y + dy[d]
	}
	return res
}

// @lc code=end

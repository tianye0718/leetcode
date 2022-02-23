/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (77.58%)
 * Likes:    599
 * Dislikes: 0
 * Total Accepted:    161.1K
 * Total Submissions: 207.8K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：[[1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 *
 *
 */
package main

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	x, y, d := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		res[x][y] = i
		a, b := x+dx[d], y+dy[d]
		if a < 0 || a >= n || b < 0 || b >= n || res[a][b] != 0 {
			d = (d + 1) % 4
		}
		x = x + dx[d]
		y = y + dy[d]
	}
	return res
}

// @lc code=end

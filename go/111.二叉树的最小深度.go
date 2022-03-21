/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (49.74%)
 * Likes:    693
 * Dislikes: 0
 * Total Accepted:    349.5K
 * Total Submissions: 701.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000
 *
 *
 */
package leetcode0111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
////////////////DFS////////////////////
// func minDepth(root *TreeNode) (depth int) {
// 	if root == nil {
// 		return 0
// 	}
// 	left := minDepth(root.Left)
// 	right := minDepth(root.Right)
// 	if left == 0 {
// 		return right + 1
// 	}
// 	if right == 0 {
// 		return left + 1
// 	}
// 	return min(left, right) + 1
// }
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

///////////////////BFS//////////////////
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
}

// @lc code=end

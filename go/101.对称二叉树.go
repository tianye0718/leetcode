/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (56.87%)
 * Likes:    1730
 * Dislikes: 0
 * Total Accepted:    488.5K
 * Total Submissions: 857.5K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给你一个二叉树的根节点 root ， 检查它是否轴对称。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,2,3,4,4,3]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,null,3,null,3]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [1, 1000] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
 *
 */
package main

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
// Recursive
// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	var helper func(left, right *TreeNode) bool
// 	helper = func(left, right *TreeNode) bool {
// 		if left == nil && right == nil {
// 			return true
// 		} else if left == nil || right == nil || left.Val != right.Val {
// 			return false
// 		} else {
// 			return helper(left.Left, right.Right) && helper(left.Right, right.Left)
// 		}
// 	}
// 	return helper(root.Left, root.Right)
// }

// Iterative
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	u := root.Left
	v := root.Right
	var queue []*TreeNode
	queue = append(queue, u)
	queue = append(queue, v)
	for len(queue) > 0 {
		u = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		v = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if u == nil && v == nil {
			continue
		} else if u == nil || v == nil || u.Val != v.Val {
			return false
		} else {
			queue = append(queue, u.Left)
			queue = append(queue, v.Right)
			queue = append(queue, u.Right)
			queue = append(queue, v.Left)
		}
	}
	return true
}

// @lc code=end

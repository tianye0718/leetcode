/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (75.31%)
 * Likes:    751
 * Dislikes: 0
 * Total Accepted:    360.2K
 * Total Submissions: 477.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[3,2,1]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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
// func postorderTraversal(root *TreeNode) (res []int) {
// 	var traversal func(node *TreeNode)
// 	traversal = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		traversal(node.Left)
// 		traversal(node.Right)
// 		res = append(res, node.Val)
// 	}
// 	traversal(root)
// 	return res
// }

// Iterative
func postorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prevNode *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prevNode {
			res = append(res, node.Val)
			prevNode = node
			root = nil
		} else {
			stack = append(stack, node)
			root = node.Right
		}
	}
	return res
}

// @lc code=end

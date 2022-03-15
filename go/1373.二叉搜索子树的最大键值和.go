/*
 * @lc app=leetcode.cn id=1373 lang=golang
 *
 * [1373] 二叉搜索子树的最大键值和
 *
 * https://leetcode-cn.com/problems/maximum-sum-bst-in-binary-tree/description/
 *
 * algorithms
 * Hard (41.26%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    11.9K
 * Total Submissions: 28.8K
 * Testcase Example:  '[1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]'
 *
 * 给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
 *
 * 二叉搜索树的定义如下：
 *
 *
 * 任意节点的左子树中的键值都 小于 此节点的键值。
 * 任意节点的右子树中的键值都 大于 此节点的键值。
 * 任意节点的左子树和右子树都是二叉搜索树。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
 * 输出：20
 * 解释：键值为 3 的子树是和最大的二叉搜索树。
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [4,3,null,1,2]
 * 输出：2
 * 解释：键值为 2 的单节点子树是和最大的二叉搜索树。
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [-4,-2,-5]
 * 输出：0
 * 解释：所有节点键值都为负数，和最大的二叉搜索树为空。
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [2,1,3]
 * 输出：6
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [5,4,8,3,null,6,3]
 * 输出：7
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树有 1 到 40000 个节点。
 * 每个节点的键值在 [-4 * 10^4 , 4 * 10^4] 之间。
 *
 *
 */
package leetcode1373

import "math"

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
type Node struct {
	IsBST bool
	Min   int
	Max   int
	Sum   int
}

var maxSum int

func maxSumBST(root *TreeNode) int {
	maxSum = 0
	traversal(root)
	return maxSum
}
func traversal(root *TreeNode) *Node {
	if root == nil {
		return &Node{
			IsBST: true,
			Min:   math.MaxInt32,
			Max:   math.MinInt32,
			Sum:   0,
		}
	}
	left := traversal(root.Left)
	right := traversal(root.Right)
	if left.IsBST && right.IsBST &&
		left.Max < root.Val && right.Min > root.Val {
		isBST := true
		min := getMin(left.Min, root.Val)
		max := getMax(right.Max, root.Val)
		sum := left.Sum + right.Sum + root.Val
		if sum > maxSum {
			maxSum = sum
		}
		return &Node{
			IsBST: isBST,
			Min:   min,
			Max:   max,
			Sum:   sum,
		}
	} else {
		return &Node{IsBST: false}
	}
}
func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

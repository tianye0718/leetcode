/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (71.35%)
 * Likes:    1147
 * Dislikes: 0
 * Total Accepted:    120.1K
 * Total Submissions: 168.2K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
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
 *
 *
 */
package leetcode0095

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
var memo [][]([]*TreeNode)

func generateTrees(n int) []*TreeNode {
	memo = make([][]([]*TreeNode), n+1)
	for i := range memo {
		memo[i] = make([]([]*TreeNode), n+1)
	}
	return build(1, n)
}
func build(lo int, hi int) (res []*TreeNode) {
	if lo > hi {
		return append(res, nil)
	}
	if memo[lo][hi] != nil {
		return memo[lo][hi]
	}
	for i := lo; i <= hi; i++ {
		leftList := build(lo, i-1)
		rightList := build(i+1, hi)
		for _, left := range leftList {
			for _, right := range rightList {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				res = append(res, root)
			}
		}
	}
	memo[lo][hi] = res
	return res
}

// @lc code=end

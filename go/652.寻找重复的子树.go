/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (57.40%)
 * Likes:    374
 * Dislikes: 0
 * Total Accepted:    42.4K
 * Total Submissions: 74K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树 root，返回所有重复的子树。
 *
 * 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 *
 * 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：root = [1,2,3,4,null,2,4,null,null,4]
 * 输出：[[2,4],[4]]
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：root = [2,1,1]
 * 输出：[[1]]
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：root = [2,2,2,3,null,3,null]
 * 输出：[[2,3],[3]]
 *
 *
 *
 * 提示：
 *
 *
 * 树中的结点数在[1,10^4]范围内。
 * -200 <= Node.val <= 200
 *
 *
 */
package main

import (
	"strconv"
)

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
var serials map[string]int
var res []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	serials = make(map[string]int, 0)
	res = make([]*TreeNode, 0)
	traversal(root)
	return res
}

func traversal(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := traversal(root.Left)
	right := traversal(root.Right)
	serial := left + "," + right + "," + strconv.Itoa(root.Val)
	serials[serial]++
	if serials[serial] == 2 {
		res = append(res, root)
	}
	return serial
}

// func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
// 	hashAll := map[string]int{}
// 	duplicate := []*TreeNode{}
// 	dfs(root, hashAll, &duplicate)
// 	return duplicate
// }

// func dfs(node *TreeNode, hashAll map[string]int, duplicate *[]*TreeNode) string {
// 	if node == nil {
// 		return "#"
// 	}
// 	lString := dfs(node.Left, hashAll, duplicate)
// 	rString := dfs(node.Right, hashAll, duplicate)
// 	//buildString := fmt.Sprintf("(%s)(%s)(%v)", lString, rString, node.Val)
// 	buildString := lString + "," + rString + "," + strconv.Itoa(node.Val)
// 	hashAll[buildString]++
// 	if hashAll[buildString] == 2 {
// 		*duplicate = append(*duplicate, node)
// 	}
// 	return buildString
// }

// @lc code=end

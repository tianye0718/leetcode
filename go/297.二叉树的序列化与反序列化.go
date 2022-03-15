/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (56.92%)
 * Likes:    784
 * Dislikes: 0
 * Total Accepted:    134.9K
 * Total Submissions: 236.9K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 *
 * 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
 *
 * 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 /
 * 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
 *
 * 提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode
 * 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,2,3,null,null,4,5]
 * 输出：[1,2,3,null,null,4,5]
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
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中结点数在范围 [0, 10^4] 内
 * -1000
 *
 *
 */
package leetcode0297

import (
	"strconv"
	"strings"
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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

const NIL = "#"
const SEP = ","

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var sb strings.Builder
	var doSerialize func(node *TreeNode)
	doSerialize = func(node *TreeNode) {
		if node == nil {
			sb.WriteString(NIL)
			sb.WriteString(SEP)
			return
		}
		val := strconv.Itoa(node.Val)
		sb.WriteString(val)
		sb.WriteString(SEP)
		doSerialize(node.Left)
		doSerialize(node.Right)
	}
	doSerialize(root)
	return sb.String()
}

// Deserializes your encoded data to tree.

func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, SEP)
	var doDeserialize func() *TreeNode
	doDeserialize = func() *TreeNode {
		if len(s) == 0 {
			return nil
		}
		if s[0] == NIL {
			s = s[1:]
			return nil
		}
		val, _ := strconv.Atoi(s[0])
		root := &TreeNode{Val: val}
		s = s[1:]
		root.Left = doDeserialize()
		root.Right = doDeserialize()
		return root
	}
	return doDeserialize()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

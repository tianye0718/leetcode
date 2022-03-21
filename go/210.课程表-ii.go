/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 *
 * https://leetcode-cn.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (55.13%)
 * Likes:    585
 * Dislikes: 0
 * Total Accepted:    113.7K
 * Total Submissions: 206.1K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中
 * prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
 *
 *
 * 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
 *
 *
 * 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：[0,1]
 * 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
 * 输出：[0,2,1,3]
 * 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
 * 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
 *
 * 示例 3：
 *
 *
 * 输入：numCourses = 1, prerequisites = []
 * 输出：[0]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= numCourses * (numCourses - 1)
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * ai != bi
 * 所有[ai, bi] 互不相同
 *
 *
 */
package leetcode0210

// @lc code=start
////////////////////////////////////DFS////////////////////////////////
// var onPath []bool
// var visited []bool
// var hasCycle bool
// var postorder []int

// func findOrder(numCourses int, prerequisites [][]int) []int {
// 	hasCycle = false
// 	visited, onPath = make([]bool, numCourses), make([]bool, numCourses)
// 	postorder = []int{}
// 	graph := buildGraph(numCourses, prerequisites)
// 	for i := 0; i < numCourses; i++ {
// 		traversal(graph, i)
// 	}
// 	if hasCycle {
// 		return []int{}
// 	}
// 	reverse(postorder)
// 	return postorder
// }
// func buildGraph(numCourses int, prerequisites [][]int) [][]int {
// 	graph := make([][]int, numCourses)
// 	for i := range graph {
// 		graph[i] = make([]int, 0)
// 	}
// 	for _, edge := range prerequisites {
// 		graph[edge[1]] = append(graph[edge[1]], edge[0])
// 	}
// 	return graph
// }
// func traversal(graph [][]int, s int) {
// 	if onPath[s] {
// 		hasCycle = true
// 		return
// 	}
// 	if visited[s] {
// 		return
// 	}
// 	visited[s] = true
// 	onPath[s] = true
// 	for _, v := range graph[s] {
// 		traversal(graph, v)
// 	}
// 	postorder = append(postorder, s)
// 	onPath[s] = false
// }
// func reverse(input []int) {
// 	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
// 		input[i], input[j] = input[j], input[i]
// 	}
// }

/////////////////////////////////////BFS/////////////////////////
func buildGraph(numCourses int, prerequisites [][]int) ([][]int, []int) {
	graph := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		indegrees[edge[0]]++
	}
	return graph, indegrees
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	// #1 build graph
	graph, indegrees := buildGraph(numCourses, prerequisites)
	// #2 push 0 indegree node to queue
	q := []int{}
	for i, v := range indegrees {
		if v == 0 {
			q = append(q, i)
		}
	}
	// traversal and record the order of traversal
	count := 0
	orders := []int{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		count++
		orders = append(orders, cur)
		for _, v := range graph[cur] {
			indegrees[v]--
			if indegrees[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if count != numCourses {
		return []int{}
	}
	return orders
}

// @lc code=end

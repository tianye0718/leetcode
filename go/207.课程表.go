/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (53.99%)
 * Likes:    1172
 * Dislikes: 0
 * Total Accepted:    183.5K
 * Total Submissions: 340K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * prerequisites[i].length == 2
 * 0 i, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */
package leetcode0207

// @lc code=start
///////////////////DFS///////////////////////////
// var visited []bool
// var onPath []bool
// var hasCycle bool

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	graph := buildGraph(numCourses, prerequisites)
// 	visited = make([]bool, numCourses)
// 	onPath = make([]bool, numCourses)
// 	hasCycle = false
// 	for i := 0; i < numCourses; i++ {
// 		traversal(graph, i)
// 	}
// 	return !hasCycle
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
// 	onPath[s] = false
// }

// func buildGraph(numCourses int, prerequisites [][]int) [][]int {
// 	graph := make([][]int, numCourses)
// 	for i := range graph {
// 		graph[i] = []int{}
// 	}
// 	for _, edge := range prerequisites {
// 		graph[edge[1]] = append(graph[edge[1]], edge[0])
// 	}
// 	return graph
// }

///////////////////////////BFS/////////////////////////
func buildGraph(numCourses int, prerequisites [][]int) ([][]int, []int) {
	graph := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		indegrees[edge[0]]++
	}
	return graph, indegrees
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	// build graph and indegree
	graph, indegrees := buildGraph(numCourses, prerequisites)
	// push node who is 0 indegree
	q := []int{}
	for i, v := range indegrees {
		if v == 0 {
			q = append(q, i)
		}
	}
	// if all nodes can be accessed, it doesn't have cycle
	count := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		count++
		for _, node := range graph[cur] {
			indegrees[node]--
			if indegrees[node] == 0 {
				q = append(q, node)
			}
		}
	}
	return count == numCourses
}

// @lc code=end

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode-cn.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (55.62%)
 * Likes:    494
 * Dislikes: 0
 * Total Accepted:    307.1K
 * Total Submissions: 552.1K
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */
package leetcode0912

// import (
// 	"math/rand"
// 	"time"
// )

// @lc code=start
// Quick Sort
// func sortArray(nums []int) []int {
// 	sz := len(nums)
// 	if sz <= 1 {
// 		return nums
// 	}
// 	last := sz - 1
// 	// randomly choose a num as a flag
// 	rand.Seed(time.Now().UnixNano())
// 	flagIndex := rand.Intn(sz)
// 	// swap flag and last num
// 	nums[flagIndex], nums[last] = nums[last], nums[flagIndex]

// 	zoneIndex := -1
// 	for i := 0; i < sz; i++ {
// 		if nums[i] <= nums[last] {
// 			zoneIndex++
// 			if i > zoneIndex {
// 				nums[i], nums[zoneIndex] = nums[zoneIndex], nums[i]
// 			}
// 		}
// 	}
// 	sortArray(nums[:zoneIndex])
// 	sortArray(nums[zoneIndex+1:])
// 	return nums
// }

// Heap Sort
// var length int
// func sortArray(nums []int) []int {
// 	length = len(nums)
// 	if length <= 1 {
// 		return nums
// 	}
// 	// #1. build the max heap
// 	buildMaxHeap(nums)
// 	// #2. swap the first item of the heap and the last item of the unsorted array
// 	// and then adjust the heap, make it be the max heap again
// 	for length > 0 {
// 		nums[0], nums[length-1] = nums[length-1], nums[0]
// 		length--
// 		adjustHeap(nums, 0)
// 	}
// 	return nums
// }
// func buildMaxHeap(nums []int) {
// 	for i := length>>1 - 1; i >= 0; i-- {
// 		adjustHeap(nums, i)
// 	}
// }
// func adjustHeap(nums []int, start int) {
// 	maxIndex := start
// 	left := start<<1 + 1
// 	right := start<<1 + 2
// 	if left < length && nums[left] > nums[start] {
// 		maxIndex = left
// 	}
// 	if right < length && nums[right] > nums[left] && nums[right] > nums[start] {
// 		maxIndex = right
// 	}
// 	if maxIndex != start {
// 		nums[start], nums[maxIndex] = nums[maxIndex], nums[start]
// 		adjustHeap(nums, maxIndex)
// 	}
// }

// Bucket Sort
// func sortArray(nums []int) []int {
// 	sz := len(nums)
// 	if sz <= 1 {
// 		return nums
// 	}
// 	min, max := nums[0], nums[0]
// 	for i := 1; i < sz; i++ {
// 		if nums[i] < min {
// 			min = nums[i]
// 		}
// 		if nums[i] > max {
// 			max = nums[i]
// 		}
// 	}
// 	bucketSize := max - min + 1
// 	buckets := make([]int, bucketSize)

// 	for _, num := range nums {
// 		buckets[num-min]++
// 	}

// 	i := 0
// 	for j, v := range buckets {
// 		if v != 0 {
// 			for k := 0; k < v; k++ {
// 				nums[i] = j + min
// 				i++
// 			}
// 		}
// 	}
// 	return nums
// }

// Merge Sort
func sortArray(nums []int) (res []int) {
	sz := len(nums)
	if sz <= 1 {
		return nums
	}
	mid := sz >> 1
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i == len(left) {
			res = append(res, right[j:]...)
			break
		}
		if j == len(right) {
			res = append(res, left[i:]...)
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	return res
}

// @lc code=end

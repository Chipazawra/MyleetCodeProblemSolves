//https://leetcode.com/problems/merge-intervals/submissions/
package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {

	resultSlice := make([][]int, 0)

	if len(intervals) == 0 {
		return resultSlice
	} else if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][len(intervals[0])-1]
	for idx := 1; idx < len(intervals); idx++ {
		if intervals[idx][0] <= end {
			end = max(intervals[idx][1], end)
		} else {
			resultSlice = append(resultSlice, []int{start, end})
			start, end = intervals[idx][0], intervals[idx][1]
		}
	}

	resultSlice = append(resultSlice, []int{start, end})

	return resultSlice
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Merge(intervals [][]int) [][]int {
	return merge(intervals)
}

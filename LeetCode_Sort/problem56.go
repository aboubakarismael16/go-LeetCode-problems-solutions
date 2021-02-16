package LeetCode_Sort

import "sort"

func merge(intervals [][]int) [][]int {
	// corner condition
	res := make([][]int,0)
	if intervals == nil || len(intervals) == 0 {
		return res
	}

	//sort by the start
	sort.Slice(intervals,func (a,b int) bool {return intervals[a][0] < intervals[b][0]})

	//add first item into res
	res = append(res,intervals[0])
	//for each compare
	for i := 1; i < len(intervals); i++ {
		curStart := intervals[i][0]
		curEnd := intervals[i][1]

		preEnd := intervals[i-1][1]
		if preEnd >= curStart {
			intervals[i-1][1] = max(preEnd, curEnd)
			intervals[i] = intervals[i-1]
		} else {
			res = append(res,intervals[i])
		}
	}

	//return
	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
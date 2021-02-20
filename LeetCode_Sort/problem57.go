package LeetCode_Sort

func intersect(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int,0)
	if intervals == nil || len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	start,end := newInterval[0], newInterval[1]
	left, right := make([][]int, 0), make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		if curEnd < start {
			left = append(left, intervals[i])
		} else if curStart > end  {
			right = append(right, intervals[i])
		} else {
			start = min(start, curStart)
			end = max(end, curEnd)
		}
	}

	temp := []int{start, end}
	res = append(res, left...)
	res = append(res, temp)
	res = append(res, right...)

	return res
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func max(a,b int) int {
    if a > b{
        return a
    } else {
        return b
    }
}
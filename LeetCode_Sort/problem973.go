package LeetCode_Sort

import "sort"

func KClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i,j int) bool {
		return distance(points[i]) < distance(points[j])
	})

	return points[:K]
}

func distance(point []int) int {
	return point[0] * point[0] + point[1] * point[1]
}
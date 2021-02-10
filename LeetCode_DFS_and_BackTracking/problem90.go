package LeetCode_DFS_and_BackTracking

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	c,res := []int{},[][]int{}
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums,k,0,c,&res)
	}

	return res
}

func generateSubsetsWithDup(nums []int,k, start int, c []int, res *[][]int)  {
	if len(c) == k {
		b := make([]int,len(c))
		copy(b,c)
		*res = append(*res,b)
		return
	}
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		c = append(c,nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}

	return
}
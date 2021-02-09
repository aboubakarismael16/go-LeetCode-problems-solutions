package LeetCode_DFS_and_BackTracking

func subsets(nums []int) [][]int {
	c, res := []int{},[][]int{}
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, c, &res)
	}

	return res
}

func generateSubsets(nums []int, k, start int, c []int, res *[][]int)  {
	if len(c) == k {
		b := make([]int,len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}

	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}

	return
}
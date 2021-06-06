package leetcode

func singleNumber(nums []int)  int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	if i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	return res
}

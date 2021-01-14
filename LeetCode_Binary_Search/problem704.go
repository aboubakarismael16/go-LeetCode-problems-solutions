package LeetCode

func search(nums []int,target int) int  {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right - left) /2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > left {
			right = mid -1
		} else {
			left = mid + 1
		}
	}

	return -1
}
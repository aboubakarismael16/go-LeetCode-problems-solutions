package LeetCode_Binary_Search

func searchInsert(nums []int,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left,right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid -1
		}
	}

	return left
}
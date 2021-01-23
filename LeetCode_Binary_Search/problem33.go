package LeetCode_Binary_Search


func search(nums []int,target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target >= nums[0] {
			if nums[mid] < target && nums[0] <= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1 
			}
		} else if target < nums[0] {
			if nums[mid] < target || nums[0] <= nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
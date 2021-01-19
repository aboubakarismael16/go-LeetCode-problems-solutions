func searchRange(nums []int,target int) []int {
	
	return []int{searchRangeFirst(nums,target),searchRangeLast(nums,target)}
}

func searchRangeFirst(nums []int,target int) int {
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchRangeLast(nums []int,target int) int {
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
		} else {
			right = mid - 1
		}
	}
	return -1
}
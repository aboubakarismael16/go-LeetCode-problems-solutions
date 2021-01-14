package LeetCode

func search(nums []int,target int) bool  {
	if nums == nil || len(nums) == 0 {
		return false
	}

	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums [mid] {
			left++
		} else if nums[left] < nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid +1
			} else {
				left = mid +1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid +1
			} else {
				right = mid -1
			}
		}
	}

	return false
}
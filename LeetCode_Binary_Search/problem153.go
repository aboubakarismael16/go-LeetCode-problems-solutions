package LeetCode_Binary_Search

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	left,right := 0,len(nums)-1
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid +1
		}
	}

	return nums[left]
}
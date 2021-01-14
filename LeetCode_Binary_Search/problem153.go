package LeetCode_Binary_Search


//first method
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

//second method

func findMin2(nums []int) int {
	min := nums[0]
	for _,num := range nums[1:] {
		if min > num {
			min = num
		}
	}

	return min
}
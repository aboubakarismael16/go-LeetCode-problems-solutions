package leetcode

// Runtime: 3584 ms
// Memory Usage: 6.8 MB

func sortArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}

		temp := nums[min_index]
		nums[min_index] = nums[i]
		nums[i] = temp
	}

	return nums
}

//second method
func quick(nums []int, start, end int) {
	if start < end {
		j, i := end, start+1

		for i <= j {
			if nums[i] <= nums[start] {
				i++
			} else if nums[j] >= nums[start] {
				j--
			} else {
				nums[j], nums[i] = nums[i], nums[j]
				j--
				i++
			}
		}

		nums[j], nums[start] = nums[start], nums[j]
		quick(nums, start, j-1)
		quick(nums, j+1, end)
	}

}

func sortArray2(nums []int) []int {
	quick(nums, 0, len(nums)-1)

	return nums
}
func quick(nums []int, start, end int) {
	if start < end {
		j, i := end, start+1

		for i <= j {
			if nums[i] <= nums[start] {
				i++
			} else if nums[j] >= nums[start] {
				j--
			} else {
				nums[j], nums[i] = nums[i], nums[j]
				j--
				i++
			}
		}

		nums[j], nums[start] = nums[start], nums[j]
		quick(nums, start, j-1)
		quick(nums, j+1, end)
	}

}

func sortArray2(nums []int) []int {
	quick(nums, 0, len(nums)-1)

	return nums
}

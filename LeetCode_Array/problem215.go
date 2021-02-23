package LeetCode_Array

import "sort"

//first mehtod
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//second method
func findKthLargest(nums []int, k int) int {
    if len(nums) == 0 {
        return 0
    }

    return selection(nums, 0, len(nums)-1, len(nums)-k)
}

func selection(nums []int, l,r,k int) int {
    if l == r {
        return nums[l]
    }

    p := partition(nums, l, r)
    if k == p {
        return nums[p]
    } else if k < p {
        return selection(nums, l, p-1, k)
    } else {
        return selection(nums, p+1, r, k)
    }
}

func partition(nums []int, low,high int) int {
    pivot := nums[high]
    i := low - 1
    for j := low; j < high; j++ {
        if nums[j] < pivot {
            i++
            nums[j], nums[i] = nums[i], nums[j]
        }
    }
    nums[i+1], nums[high] = nums[high], nums[i+1]

    return i + 1
}
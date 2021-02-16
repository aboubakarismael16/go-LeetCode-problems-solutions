package LeetCode_Sort

import (
	"sort"
	"strconv"
)


func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	var number string
	sort.Slice(nums, func (i,j int) bool {
		first,second := strconv.Itoa(nums[i]),strconv.Itoa(nums[j])

		if first+second >= second+first {
			return true
		} else {
			return false
		}
	})

	for _, num := range nums {
		number += strconv.Itoa(num)
	}
	if number[0] == '0' {
		return number[:1]
	}

	return number
}
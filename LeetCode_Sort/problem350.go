package LeetCode_Sort

import "sort"

//first method
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	i,j := 0,0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		}
		if i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j] {
			i++
		}
		if i < len(nums1) && j < len(nums2) && nums2[j] < nums1[i] {
			j++
		}
	}

	return res
}

//second method

func intersect2(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return res
	}

	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if val,ok := m[nums2[i]]; ok {
			if val > 0 {
				res = append(res, nums2[i])
			}
			m[nums2[i]]--
		}
	}

	return res
}
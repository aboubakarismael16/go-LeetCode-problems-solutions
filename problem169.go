package LeetCode

LeetCode problem169

// first method
func majorityElement(nums []int) int()  {
	if len(nums) == 1 {
		return nums[0]
	}	

	thresold := 1 + len(nums)/2
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val,ok := m[nums[i]]
		if ok {
			val++
			if val >= thresold {
				return m[i]
			} else {
				m[nums[i]] = val
			}
		} else {
			m[nums[i]] = 1
		}
	}

	return 0
}

//second method
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	for _,v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}

	return 0
}
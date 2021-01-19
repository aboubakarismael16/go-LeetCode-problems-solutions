package LeetCode_Array


func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	threshold := 1 + len(nums)/2
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		val,ok := m[nums[i]]
		if ok {
			val++
			if val >= threshold {
				return nums[i]
			} else {
				m[nums[i]] = val
			}
		} else {
			m[nums[i]] = 1
		}
	}
	return 0
}
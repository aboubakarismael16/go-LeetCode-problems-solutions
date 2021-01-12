package LeetCode

func getSum(a,b int) int  {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	return getSum((a&b)<<1,a^b)
}
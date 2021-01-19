package LeetCode_Array



func plusOne(digits []int) []int {
	res := make([]int,0)
	if digits == nil || len(digits) == 0 {
		return res
	}

	carry := 1
	for i := len(digits)-1; i >= 0; i-- {
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10 
	}
	if carry == 1 {
		res = []int{1}
		digits = append(res,digits...)
	}

	return digits
}
package LeetCode

func isPowerOfTwo(n int) bool {
	for n >= 2 {
		if n%2 == 0 {
			n = n/2 
		} else {
			return false
		}
	}

	return n == 1
}
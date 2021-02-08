package LeetCode

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1/x
	}
	temp := myPow(x,n/2)
	if n % 2 == 0 {
		return temp * temp
	}

	return temp * temp * x
}
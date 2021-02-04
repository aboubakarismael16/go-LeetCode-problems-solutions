package LeetCode

func reverse(x int) int {
	if x == 0  {
		return 0
	}

	temp := 0
	for x != 0 {
		temp = temp * 10 + (x % 10)
		x = x / 10
	}

	if temp < -(1<<31) || temp > 1<<31-1 {
		return 0
	}

	return temp
}
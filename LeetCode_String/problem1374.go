package LeetCode_String

func generateTheString(n int) string {
	if n == 0 {
		return ""
	}
	if n % 2 == 0 {
		// pair
		return buildString(n-1) + "b"
	} else {
		//impair
		return buildString(n)
	}
}

func buildString(n int) string  {
	res := ""
	for i := 0; i < n ; i++ {
		res += "a"
	}

	return res 
}
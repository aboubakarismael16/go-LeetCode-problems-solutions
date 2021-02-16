package LeetCode_Sort

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sByte,tByte := []byte(s),[]byte(t)
	m := make(map[byte]int)
	for i := 0; i < len(sByte); i++ {
		m[sByte[i]]++
	}
	for i := 0; i < len(tByte); i++ {
		m[tByte[i]]--
		if m[tByte[i]] < 0 {
			return false
		}
	}

	return true
}
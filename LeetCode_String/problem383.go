package LeetCode_String

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	r := []byte(ransomNote)
	m := []byte(magazine)

	dic := make(map[byte]int)
	for i := 0; i < len(m); i++ {
		dic[m[i]]++
	}

	for i := 0; i < len(r); i++ {
		val,ok := dic[r[i]]
		if ok {
			if val == 0 {
				return false 
			} else {
				dic[r[i]]--
			}
		} else {
			return false
		}
	}

	return true
}
package LeetCode_String

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	res := strs[0]
	for i := 1; i < len(strs); i++ {
		resLenght := len(res)
		tempLength := len(strs[i])
		length := min(resLenght,tempLength)
		j := 0
		for j < length {
			if res[j] != strs[i][j] {
				break
			}
			j++
		}
		res = res[:j]
		if len(res) == 0 {
			return ""
		}
	}
	
	return res
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
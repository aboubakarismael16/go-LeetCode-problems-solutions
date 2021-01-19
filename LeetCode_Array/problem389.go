func findTheDifference(s string,t string) byte {
	var res byte
	if len(s) >= len(t) {
		return res
	}
	S := []byte(s)
	T := []byte(t)

	dic := make(map[byte]int)
	for i := 0; i <len(S); i++ {
		dic[S[i]]++
	}
	for i := 0; i < len(T); i++ {
		val,ok := dic[T[i]]
		if !ok {
			res = T[i]
			break
		} else {
			if val == 0 {
				res = T[i]
				break
			}
			dic[T[i]]--
		}
	}

	return res
}
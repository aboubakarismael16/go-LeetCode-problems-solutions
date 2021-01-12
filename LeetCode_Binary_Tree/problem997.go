package LeetCode_Binary_Tree

func findJudge(N int,trust [][]int) int {
	if N == 1 {
		return N
	}

	dic := make(map[int]int)
	for _,v := range trust {
		dic[v[1]]++
	}

	res := -1
	for k,v := range dic {
		if v == N-1 {
			res  = k
			break
		}
	}

	if v > -1 {
		for _,v := range trust {
			if v[0] == res {
				return -1 
			}
		}
	}

	return res
}
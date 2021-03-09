package LeetCode

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}

	visited := make([]bool, len(isConnected))
	res := 0
	for i := range isConnected {
		if !visited[i] {
			dfs(isConnected, i, visited)
			res++
		}
	}

	return res
}

func dfs(isConnected [][]int, cur int, visited []bool)  {
	visited[cur] = true
	for j := 0; j < len(isConnected[cur]); j++ {
		if !visited[j] && isConnected[cur][j] == 1 {
			dfs(isConnected, j, visited)
		}
	}
}
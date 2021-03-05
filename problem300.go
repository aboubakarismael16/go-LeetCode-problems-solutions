package LeetCode

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	res := 0

	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] < nums[j-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i]+1
		res = max(res, dp[i])
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
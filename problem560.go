package LeetCode

func subarraySum(nums []int, k int) int {
	ret, summary := 0, 0
	magic := make(map[int]int) 
    magic[0] = 1
	for _, v := range nums {
        summary += v
        ret += magic[summary-k]
        magic[summary]++
	}
	return ret
}
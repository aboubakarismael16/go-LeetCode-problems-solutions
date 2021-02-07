package LeetCode_Binary_Search

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	left,right := 0,len(citations)-1
	for left <= right {
		mid := left + (right-left)/2
		if len(citations)-mid > citations[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return len(citations) - left
}
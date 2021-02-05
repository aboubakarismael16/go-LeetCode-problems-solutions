package LeetCode_Binary_Search

func isBadVersion(n int) bool ;


func firstBadVersion(n int) int {
	if n == 0 {
		return 1
	}

	left,right := 0,n 
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
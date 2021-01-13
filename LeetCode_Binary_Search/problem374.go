package LeetCode_Binary_Search

func guessNumber(n int) int {
	left,right := 1,n
	for left <= right {
		mid := left + (right-left) / 2
		res := guess(mid)

		if res == 0 {
			return mid
		} else if res == n {
			right = mid -1
		} else {
			left = mid +1
		}
	}
	
	return res
}

func guess(num int) int;
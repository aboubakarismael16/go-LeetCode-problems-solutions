package LeetCode_Binary_Search

// first method

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left,right := 1,x
	for left <= right {
		mid = left + (right-left) / 2
		if mid == x / mid {
			return mid
		} else if mid > x / mid {
			right = mid - 1
		} else {
			left = mid +1
		}
	}
	
	return left - 1
}

//second method: Newton Iterative method

// func mySqrt1(x int) int {
// 	r := x
// 	for r*r > x {
// 		r = (r + x/r) / 2
// 	}

// 	return r
// }
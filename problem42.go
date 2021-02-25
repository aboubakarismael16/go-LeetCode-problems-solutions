package LeetCode

func trap(height []int) int {
	res := 0
    left, right := 0, len(height)-1
    maxLeft,maxRight := 0,0
    
    for left <= right {
        if height[left] <= height[right] {
            if height[left] > maxLeft {
                maxLeft = height[left]
            } else {
                res += maxLeft - height[left] 
            }
            left++
        } else {
            if height[right] >= maxRight {
                maxRight = height[right]
            } else {
                res += maxRight - height[right]
            }
            right--
            }
        }
    
    return res
}
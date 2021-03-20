package leetcode

func longestConsecutive(nums []int) int {
    res , numMap := 0, map[int]int{}
    for _, num := range nums {
        if numMap[num] == 0 {
            left, right, sum := 0, 0, 0
            if numMap[num - 1] > 0 {
                left = numMap[num - 1]
            } else {
                left = 0
            }
            if numMap[num + 1] > 0 {
                right = numMap[num + 1]
            } else {
                right = 0
            }
            sum = left + right + 1
            numMap[num] = sum
            res = max(res, sum)
            numMap[num -left] = sum
            numMap[num + right] = sum
        } else {
            continue
        }
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
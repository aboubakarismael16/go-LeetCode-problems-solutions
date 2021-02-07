package LeetCode_Binary_Search

// first method
func findDuplicate(nums []int) int {
    left,right := 0,len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        count := 0
        for _,num := range nums {
            if num <= mid {
                count++
            }
        }
        if count > mid {
            right = mid
        } else {
            left = mid + 1
        }
    }
    
    return left
}


// second method

// func findDuplicate(nums []int) int {
    
//     slow,fast := nums[0],nums[nums[0]]
//     for fast != slow {
//         slow = nums[slow]
//         fast = nums[nums[fast]]
//     }
    
//     walker := 0
//     for walker != slow {
//         walker = nums[walker]
//         slow = nums[slow]
//     }
    
//     return walker
// }

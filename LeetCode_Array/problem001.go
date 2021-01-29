package LeetCode_Array


func TwoSum(nums []int,target int) []int  {
	res :=[]int{-1,-1}

	lookup := make(map[int]int)
	for i, num := range nums {
		temp := target - num 
		if _,ok := lookup[temp]; ok {
			res[0] = lookup[temp]
			res[1] = i
			return res
		}
		lookup[num] = i
	}
	return res
	
}
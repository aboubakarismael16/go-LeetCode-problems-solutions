# liste

## 1. liste1

[leetcode](http://leetcode.com)

``` go
func twoSum(nums []int,target int) []int {
	m := make(map[int]int)

 	for i := 0; i <len(nums)  ; i++ {
  	another := target - nums[i]

		if _,ok := m[another]; ok {
			return []int{m[another],i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{12,34,5,6,11}
	target := 11
	fmt.Println(twoSum(nums,target))

}

```


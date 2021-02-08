package LeetCode_DFS_and_BackTracking

var (
	letterMap = []string{
		" ",    //0
		"",		//1
		"abc",	//2
		"def",	//3
		"ghi",	//4
		"jkl",	//5
		"mno", 	//6
		"pqrs",	//7
		"tuv",	//8
		"wxyz",	//9
	}
	res = []string{}
	final = 0
)

// DFS method
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	res = []string{}
	findCombination(&digits,0,"")

	return res 
}

func findCombination(digits *string, index int, s string)  {
	if index == len(*digits) {
		res = append(res,s)
		return
	}

	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(digits,index+1,s+string(letter[i]))
	}

	return
}

// No-recursive method

func letterCombinations_(digits string) []string {
	if digits == "" {
		return []string{}
	}

	index := digits[0]-'0'
	letter := letterMap[index]
	temp := []string{}
	for i := 0; i < len(letter); i++ {
		if len(res) == 0 {
			res = append(res,"")
		}
		for j := 0; j < len(res); j++ {
			temp = append(temp,res[j]+string(letter[i]))
		}
	}

	res = temp
	final++
	letterCombinations(digits[1:])
	final--
	if final == 0 {
		temp = res 
		res = []string{}
	}


	return temp
}
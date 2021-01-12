package main

import "fmt"

func findTheLongstring(s string) int  {
	runes := []rune(s)
	runeMaps := make(map[rune]int)
	longest := 0
	preLenght := 0

	for i,rn := range runes {
		var lenght int
		if val,ok := runeMaps[rn]; !ok || val < i - preLenght {
			lenght = preLenght + 1
		} else {
			lenght = i - val
		}
		if lenght > longest {
			longest = lenght
		}
		preLenght = lenght
		runeMaps[rn] = i
	}

	return longest
}


func main() {
	s := "sadhashdssss"
	r := ""
	fmt.Println(findTheLongstring(s))
	fmt.Println(findTheLongstring(r))
	
}
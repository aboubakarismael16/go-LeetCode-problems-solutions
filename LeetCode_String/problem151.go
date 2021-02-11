package LeetCode_String
import "strings"

func reverseWord(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse151(m *[]string, i,j int)  {
	for i <= j {
		(*m)[i],(*m)[j] = (*m)[j],(*m)[i]
		i++
		j--
	} 
}
package LeetCode_Binary_Tree

type Node struct {
	Val int
	Children []*Node
}

func levelOrder2(root *Node) [][]int  {
	res := [][]int{}
	if root == nil {
		return res 
	}

	var q []*Node 
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}
		
		for i := 0; i < lenght; i++ {
			node := q[0]
			q = q[1:]
			level= append(level,node.Val)

			for _,v := range node.Children {
				q = append(q,v)
			}
		}
		if len(level) > 0 {
			res = append(res,level)
		}
		
		
	}


	return res	
}

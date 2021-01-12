package LeetCode_Binary_Tree


func levelOrder(root *TreeNode) [][]int  {
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0 ; i < lenght; i++ {
			node := q[0]
			q = q[1:]
			level = append(level,node.Val)

			if node.Left != nil {
				q = append(q,node.Left)
			}
			if node.Right != nil {
				q = append(q,node.Right)
			}

			for len(level) > 0 {
				res = append(res,level)
			}
		}
	}

	return res
}

// func levelOrder(root *TreeNode) [][]int {
//     levels := [][]int{}
//     dfsLevel(root,-1,&levels)
//     return levels
// }

// func dfsLevel(node *TreeNode,level int,res *[][]int) {
//     if node == nil {
//         return 
//     }
    
//     currLevel := level + 1
//     for len(*res) <= currLevel {
//         *res = append(*res,[]int{})
        
//     }
//     (*res)[currLevel] = append((*res)[currLevel],node.Val)
//     dfsLevel(node.Left,currLevel,res)
//     dfsLevel(node.Right,currLevel,res)
// }



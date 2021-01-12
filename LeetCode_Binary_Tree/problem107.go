package LeetCode_Binary_Tree

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	var q []*TreeNode
	q = append(q,root)

	for len(q) > 0 {
		lenght := len(q)
		level := []int{}

		for i := 0; i < lenght; i++ {
			node := q[0]
			q = q[1:]

			level = append(level,node.Val)
			if node.Left != nil {
				q = append(q,node.Left)

			}
			if node.Right != nil {
				q = append(q,node.Right)
			}
		}

		if len(level) > 0 {
			res = append(res,level)
		}
	}
	result := [][]int{}
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result,res[i])
	}
	return result
	
}

package LeetCode_Binary_Tree

func findMode(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	dic := make(map[int]int)
	inorder(root,&dic)
	max := 0
	for _,v := range dic {
		if v > max {
			max = v
		}
	}
	for k,v := range dic {
		if v == max {
			res = append(res,k)
		}
	}

	return res
}

func inorder(root *TreeNode, dic *map[int]int)  {
	if root == nil {
		return
	}

	inorder(root.Left,dic)
	(*dic)[root.Val]++
	inorder(root.Right,dic)
}
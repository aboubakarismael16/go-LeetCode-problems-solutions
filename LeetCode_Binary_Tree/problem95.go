package LeetCode_Binary_Tree


func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generateBSTree(1,n)
}

func generateBSTree(start,end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree,nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBSTree(start,i-1)
		right := generateBSTree(i+1,end)

		for _,l := range left {
			for _,r := range right {
				root := &TreeNode{
					Val : i,
					Left : l,
					Right :r,
				}
				tree = appen(tree,root)
			}
		}
	}

	return tree
}
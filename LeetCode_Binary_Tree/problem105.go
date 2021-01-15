package LeetCode_Binary_Tree

func buildTree(preorder []int,inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}	

	return buildTreeHelper(preorder,inorder)
}

func buildTreeHelper(preorder []int,inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) ==  0 {
		return nil
	}

	rootVal := preorder[0]
	index := 0
	for i,v := range inorder {
		if v == rootVal {
			index = i
		}
	}

	root :=&TreeNode{}
	root.Val = rootVal
	root.Left  = buildTreeHelper(preorder[1:index+1],inorder[:index])
	root.Right = buildTreeHelper(preorder[index+1:],inorder[index+1:])

	return root
	
}
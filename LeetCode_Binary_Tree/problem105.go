package LeetCode_Binary_Tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	return buildTreeHelper(preorder, inorder)
}

func buildTreeHelper(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	index := 0
	for i, v := range inorder {
		if v == rootVal {
			index = i
		}
	}

	root := &TreeNode{}
	root.Val = rootVal
	root.Left = buildTreeHelper(preorder[1:index+1], inorder[:index])
	root.Right = buildTreeHelper(preorder[index+1:], inorder[index+1:])

	return root

}

//second method
func buildTree2(preorder []int, inorder []int) *TreeNode {
	i := 0
	for _, val := range inorder {
		if val == preorder[0] {
			break
		}
		i++
	}
	if len(preorder) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
}

//recursion
func buildTree3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
